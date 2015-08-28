package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "flag"
    "os"
    "strings"
)

func main() {
    infile    := flag.String("i", "stdin", "O arquivo CSV de origem")
    separator := flag.String("s", ";", "O caracter que separa os campos")
    formato   := flag.String("f", "", "O formato de saida. Deve casar com o numero de campos do CSV, separar por virgula")
    pessoa    := flag.Bool("p", false, "Assume o formato \"pessoa\" -60,-80,-60,-1,-20,-20,-10")
    outfile   := flag.String("o", "stdout", "O arquivo de saida (Fixed Width Format)")
    flag.Parse()

    var out io.WriteCloser
    var in  io.ReadCloser
    var err error

    if *infile == "stdin"{
        in = os.Stdin
    }else{
        in, err = os.Open(*infile)
        if err != nil{
            log.Fatal(err)
        }
    }

    if *outfile == "stdout"{
        out = os.Stdout
    }else{
        out, err = os.Create(*outfile)
        if err != nil{
            log.Fatal(err)
        }
    }

    if *pessoa {
        *formato = "-60,-80,-60,-1,-20,-20,-10"        
    }

    fields := strings.Split(*formato, ",")
    for idx,field := range fields {
        fields[idx] = "%"+field+"s"
    }

    r := csv.NewReader(in)
    switch *separator{
    case ",":
        r.Comma = ','
    case ";":
        r.Comma = ';'
    }

    defer in.Close()
    defer out.Close()

    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        for idx,field := range fields {
            if idx < len(record) {
                fmt.Fprintf(out,field,record[idx])
            }
        }        
        fmt.Fprintf(out,"\n")
    }
}
