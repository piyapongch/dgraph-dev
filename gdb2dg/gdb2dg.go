package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("arguments error, [graphdb n-quads export file] [n-quads output file]")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	out, err := os.Create(os.Args[2])
	scanner := bufio.NewScanner(file)
	var t, s, p, o string
	for scanner.Scan() {
		s = strings.Split(scanner.Text(), " ")[0] // subject
		p = strings.Split(scanner.Text(), " ")[1] // predicate
		o = strings.Split(scanner.Text(), " ")[2] // object
		if t != s {
			// add external id edge
			out.WriteString(s + " <xid> \"" + strings.Trim(s, "<>") + "\" .\n")
			// add generic dgraph node type, IRObject
			out.WriteString(s + " <dgraph.type> \"IRObject\" .\n")
			t = s
		}
		if strings.Contains(scanner.Text(), "#long>") {
			// replace xsd:long RDF type (not supported in dgraph) with xsd:double
			ds := strings.Replace(scanner.Text(), "#long>", "#double>", 1) + "\n"
			out.WriteString(ds)
		} else {
			// convert IRI object to string for empty uid when import to dgraph
			switch p {
			case "<info:fedora/fedora-system:def/model#hasModel>":
				// add dgraph.type
				out.WriteString(s + " <dgraph.type> " + o + " .\n")
				out.WriteString(scanner.Text() + "\n")
			case "<http://www.loc.gov/premis/rdf/v1#hasMessageDigest>":
				out.WriteString(s + " " + p + " \"" + o + "\" .\n")
			case "<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>":
				out.WriteString(s + " " + p + " \"" + o + "\" .\n")
			case "<http://www.w3.org/ns/ldp#hasMemberRelation>":
				out.WriteString(s + " " + p + " \"" + o + "\" .\n")
			case "<http://www.w3.org/ns/ldp#insertedContentRelation>":
				out.WriteString(s + " " + p + " \"" + o + "\" .\n")
			case "<http://fedora.info/definitions/v4/repository#hasFixityService>":
				out.WriteString(s + " " + p + " \"" + o + "\" .\n")
			case "<http://www.iana.org/assignments/relation/describedby>":
				out.WriteString(s + " " + p + " \"" + o + "\" .\n")
			default:
				// write original triple
				out.WriteString(scanner.Text() + "\n")
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	out.Sync()
	fmt.Println("Done.")
}
