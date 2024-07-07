package main

import (
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/code-generator/examples/apiserver/apis/example"
	examplev1 "k8s.io/code-generator/examples/apiserver/apis/example/v1"
	"sigs.k8s.io/yaml"
)

func main() {
	start := metav1.NewTime(time.Date(1994, 1, 1, 0, 0, 0, 0, time.Local))
	end := metav1.NewTime(time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local))

	src := &examplev1.TestTypeStatus{
		Blah: "blah",
		Times: &examplev1.Times{
			Start: &start,
			End:   &end,
		},
	}
	dst := &example.TestTypeStatus{}
	examplev1.Convert_v1_TestTypeStatus_To_example_TestTypeStatus(src, dst, nil)

	yamlSrc, _ := yaml.Marshal(src)
	fmt.Printf("examplev1.TestTypeStatus:\n%s\n---\n", string(yamlSrc))
	yamlDst, _ := yaml.Marshal(dst)
	fmt.Printf("example.TestTypeStatus:\n%s\n---", string(yamlDst))
	fmt.Printf("conversion ok: %t", dst.Times.Start.Equal(&start) && dst.Times.End.Equal(&end))
}
