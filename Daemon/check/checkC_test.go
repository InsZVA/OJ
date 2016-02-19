//Run this test after run ./server/test.go, a simple server to response a json
package check

import "testing"
import "fmt"

func TestGetStandardInput(t *testing.T) {
    checkC := NewCheckC(2, "#include \"stdio.h\"\nint main(){printf(\"hello world!\");return 0;}")
    err := checkC.GetStandardInOut()
    if err != nil {
        t.Error(err)
    }
}

func TestBuild(t *testing.T) {
    checkC := NewCheckC(2, "#include \"stdio.h\"\nint main(){printf(\"hello world!\");return 0;}")
    err := checkC.Build()
    if err != nil {
        t.Error(err)
    }
}

func TestCheck(t *testing.T) {
    submission := `#include <stdio.h>` + "\n"
    submission += `int main() { int i; int c[100000];scanf("%d",&i); while(1){} printf("%d\n%d",i,i); return 0; }`
    checkC := NewCheckC(2, submission)
    err := checkC.Build()
    if err != nil {
        t.Error(err)
    }
    runinfo, err := checkC.Check()
    if err != nil {
        t.Error(err)
    }
    fmt.Println("Time:\t", runinfo.Time, "ms")
    fmt.Println("Memory:\t", runinfo.Memory, "kb")
    if err == nil {
        fmt.Println("Accepted!")
    }
}