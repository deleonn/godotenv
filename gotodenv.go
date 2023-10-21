package godotenv

import (
    "os"
    "strings"
)

func Config(defaultPath string) {
    path := ".env"

    if defaultPath != "" {
        path = defaultPath
    }

    contents, _ := os.ReadFile(path) 
    bytesToString := string(contents)
    splitted := strings.Split(bytesToString, "\n")

    for _, v := range splitted {
        slices := strings.Split(v, "=")
        joinedValues := strings.Join(slices[1:], "=")

        key := slices[0]

        if strings.HasPrefix(joinedValues, "\"") && strings.HasSuffix(joinedValues, "\"") {
            joinedValues = joinedValues[1:len(joinedValues)-1]
        }

        os.Setenv(key, joinedValues)
    }
}
