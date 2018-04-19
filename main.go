package main

func main() {
	// c := cli.NewCLI("Qa", "1.0.0")
	// c.Args = os.Args[1:]
	// c.Commands = map[string]cli.CommandFactory{
	// 	"foo": fooCommandFactory,
	// 	"bar": barCommandFactory,
	// }
	//
	// exitStatus, err := c.Run()
	// if err != nil {
	// 	log.Println(err)
	// }
	//
	// os.Exit(exitStatus)

	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// files := List(dir)
	// fmt.Println(files)

	// fileList := make([]string, 0)
	//
	// err = filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
	// 	if extRe.MatchString(path) {
	// 		fileList = append(fileList, path)
	// 	}
	// 	return err
	// })
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// for _, file := range fileList {
	// 	inFile, err := os.Open(file)
	//
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// 	defer inFile.Close()
	//
	// 	var partials []*partial
	// 	var section string
	//
	// 	partialRow := &partial{}
	//
	// 	scanner := bufio.NewScanner(inFile)
	// 	scanner.Split(bufio.ScanLines)
	//
	// 	for scanner.Scan() {
	// 		ln := scanner.Text()
	//
	// 		if groupRe.MatchString(ln) {
	// 			section = groupLeader
	// 		} else if actionRe.MatchString(ln) {
	// 			section = actionLeader
	// 		} else if expectRe.MatchString(ln) {
	// 			section = expectLeader
	// 		} else if breakRe.MatchString(ln) {
	// 			partials = append(partials, partialRow)
	// 			partialRow = &partial{}
	// 			continue
	// 		}
	//
	// 		str := cleanString(ln, section)
	//
	// 		if len(str) > 0 {
	// 			switch section {
	// 			case groupLeader:
	// 				partialRow.group.WriteString(str)
	// 				partialRow.group.WriteString(" ")
	// 			case actionLeader:
	// 				partialRow.action.WriteString(str)
	// 				partialRow.action.WriteString(" ")
	// 			case expectLeader:
	// 				partialRow.expect.WriteString(str)
	// 				partialRow.expect.WriteString(" ")
	// 			}
	// 		}
	// 	}
	//
	// 	partials = append(partials, partialRow)
	//
	// 	for _, p := range partials {
	// 		fmt.Println(p.group.String())
	// 		fmt.Println(p.action.String())
	// 		fmt.Println(p.expect.String())
	// 		fmt.Println("---------")
	// 	}
	// }
}
