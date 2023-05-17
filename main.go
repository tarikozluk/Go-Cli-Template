package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "Golang Cli Tutorial",
		Short: "My App is a simple CLI application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hallo Zusammen!")
		},
	}

	insertCmd := &cobra.Command{
		Use:   "insert [name] [surname]",
		Short: "Insert name surname and datetime into txt file",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			ad := args[0]
			soyad := args[1]
			tarih := time.Now().Format("02.01.2006")

			f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer f.Close()

			_, err = fmt.Fprintf(f, "%s - %s %s\n", tarih, ad, soyad)
			if err != nil {
				fmt.Println("Error on writing file:", err)
				return
			}

			fmt.Printf("%s - %s %s to txt added.\n", tarih, ad, soyad)
		},
	}

	create_file := &cobra.Command{
		Use:   "create [file] [extension]",
		Short: "a cli solution the create a file like mkdir :)",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			file_name := args[0]
			extension := args[1]
			file, err := os.Create(file_name + "." + extension)
			if err != nil {
				fmt.Println("Fatal Error:", err)
				return
			}
			defer file.Close()

			fmt.Println("File successfully created:", file_name+"."+extension)
		},
	}
	//todo: create an api and get data from there with cli commands (go run main.go last_value kpi server:port)
	rootCmd.AddCommand(create_file)
	rootCmd.AddCommand(insertCmd)

	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
