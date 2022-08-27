package env

import (
	"bufio"
	"os"
	"regexp"
)

var regexEnvs = regexp.MustCompile(`(\S+)=(\S+)`)

func DefineEnvs(filename ...string) error {
	for c := 0; c < len(filename); c++ {
		file, err := os.Open(filename[c])

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		if err != nil {
			return err
		}

		sc := bufio.NewScanner(file)

		for sc.Scan() {
			envMatch := regexEnvs.FindStringSubmatch(sc.Text())
			if envMatch != nil {
				err := os.Setenv(envMatch[1], envMatch[2])
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
