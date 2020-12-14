package k8sclient

import (
	"context"
	"encoding/json"
	"fmt"
	"glados-manager/svccache"
	"glados-manager/types"
	"log"
	"strconv"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewK8sClientset(kubeConfPath string) *kubernetes.Clientset {
	client := NewK8sClientsetV2(kubeConfPath, true)
	if client == nil {
		client = NewK8sClientsetV2(kubeConfPath, false)
	}

	return client
}

func NewK8sClientsetV2(kubeConfPath string, isSecure bool) *kubernetes.Clientset {
	var conf *rest.Config
	var err error

	conf, err = clientcmd.BuildConfigFromFlags("", kubeConfPath)
	if err != nil {
		fmt.Println("Error connecting to cluster:", err)
		return nil
	}

	if err != nil {
		panic(err.Error())
	}

	// TODO for service account configuration
	// conf, err = rest.InClusterConfig()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	// if err != nil {
	// 	panic(err.Error())
	// }

	clientset, err := kubernetes.NewForConfig(conf)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return clientset
}

//GetListOfNamespaces test k8s communication
func GetListOfNamespaces(kubeConfPath string) ([]byte, error) {
	// var namespacePodData []Poddata
	// namespacePodData, err := CalculatePodPasesPerNamespace(kubeConfPath)
	// if err != nil {
	// 	return nil, err
	// }
	namespaceslist, err := GetNamespacesViaAPI(kubeConfPath)
	if err != nil {
		return nil, err
	}
	var namespacek8s []types.Namespacedata
	for _, v := range namespaceslist.Items {
		fmt.Println(v.ObjectMeta.Name)
	}
	jsonstr, err := json.Marshal(namespacek8s)
	if err != nil {
		return nil, err
	}
	return jsonstr, nil
}

//GetNamespacesViaAPI to test that the k8s client is working
func GetNamespacesViaAPI(kubeConfPath string) (*v1.NamespaceList, error) {

	clientset := NewK8sClientset(kubeConfPath)
	namespaces, err := clientset.
		CoreV1().
		Namespaces().
		List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}
	return namespaces, nil
}

//ServiceCreate create a nodeport service
func ServiceCreate(kubeConfPath string, reqd types.SvcRequest) error {

	//check if a service already exists
	if !svccache.CheckSvcToCache(reqd) {
		return fmt.Errorf("NodePort already exists in cache")
	}
	log.Print("Creating server on nodeport " + strconv.Itoa(int(reqd.NodePort)))
	clientset := NewK8sClientset(kubeConfPath)
	_, err := clientset.CoreV1().Services(reqd.Namespace).Create(context.TODO(), &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      reqd.Label.Value + "-" + strconv.Itoa(int(reqd.NodePort)),
			Namespace: reqd.Namespace,
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Name:     reqd.Label.Value + "-" + strconv.Itoa(int(reqd.Port)),
					Protocol: "TCP",
					Port:     reqd.Port,
					NodePort: reqd.NodePort,
				},
			},
			Selector: map[string]string{
				reqd.Label.Key: reqd.Label.Value,
			},
			Type: "NodePort",
		},
	}, metav1.CreateOptions{})

	// alert on error
	if err != nil {
		log.Print("Service creation failed")
		log.Print(err)
	} else {
		// fmt.Println(reqd)
		svccache.AddToCache(reqd)
		log.Print("Nodeport service successfully created")
	}
	return nil
}

//ServiceDelete create a nodeport service
func ServiceDelete(kubeConfPath string, reqd types.SvcRequest) error {

	servicename := reqd.Label.Value + "-" + strconv.Itoa(int(reqd.NodePort))
	clientset := NewK8sClientset(kubeConfPath)
	err := clientset.CoreV1().Services(reqd.Namespace).Delete(context.TODO(), servicename, metav1.DeleteOptions{})

	// alert on error
	if err != nil {
		log.Print("Service deletion failed")
		log.Print(err)
	} else {
		// fmt.Println(reqd)
		if svccache.RemoveFromCache(reqd) {
			log.Print("Nodeport service successfully deleted")
		} else {
			log.Print("Service could not be deleted from cache")
		}
	}
	return nil
}
