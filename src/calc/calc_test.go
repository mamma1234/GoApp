package calc_test

import (
    "calc"
    "testing"
)

func TestSum(t *testing.T) {
    s := calc.Sum(1, 2, 3)

    if s != 6 {
        t.Error("Wrong result")
    }
}


/*
PS C:\github\mamma1234\GoApp\src> cd calc
PS C:\github\mamma1234\GoApp\src\calc> go test
PASS
ok      calc    0.450s
PS C:\github\mamma1234\GoApp\src\calc>
*/