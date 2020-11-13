package main

import "testing"

var (
	value            Value   = Value{922337201}
	valseharusnya 	 int = 425352956635425801
)

func TestAddUpTo(t *testing.T) {
    t.Logf("ValueAddUpTo : %d", value.AddUpTo())

    if value.AddUpTo() != valseharusnya {
        t.Errorf("SALAH! harusnya %d", valseharusnya)
    }
}

func TestAddUpToOpt(t *testing.T) {
    t.Logf("ValueAddUpToOpt : %d", value.AddUpToOpt())

    if value.AddUpToOpt() != valseharusnya {
        t.Errorf("SALAH! harusnya %d", valseharusnya)
    }
}



func BenchmarkAddUpTo(b *testing.B) {
    b.ReportAllocs()
	value.AddUpTo()
	
}

func BenchmarkAddUpToOpt(b *testing.B) {
    b.ReportAllocs()
	value.AddUpToOpt()
}

