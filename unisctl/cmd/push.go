package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var pushUsage = `Usage:  unisctl push [OPTIONS] /PATH/IAMGE[:TAG]

Options:
  -f, --configure-file  Add configure file with image
  -h, --help            help for push
  -p, --public-image    Push an image as a public one
  -t, --type            Point out type of image, like(docker or unikernel)
`

var cfgFile string
var pushPublicFlag bool //decide whether image is public or private
var imageType string    //decide whether image is docker images or unikernel image

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push an image to registry",
	Long:  "Push an image to registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//make sure the file exists
		arg := args[0]
		var path string
		var imageID string
		var size string
		if strings.Contains(arg, ":") {
			path = strings.Split(arg, ":")[0]
		} else {
			path = arg
		}
		fileInfo, err := os.Stat(path)
		if err != nil {
			logrus.Fatal(err)
		}

		//get image size
		size = strconv.FormatInt(fileInfo.Size()/1024/1024, 10)

		//generate imageID
		if content, err := ioutil.ReadFile(path); err != nil {
			logrus.Fatal(err)
		} else {
			temp := sha256.Sum256(content)
			imageID = hex.EncodeToString(temp[:])
		}

		//make sure the format is correct and get image, tag
		var image string
		var tag string
		if strings.Contains(arg, "/") {
			splitedArg := strings.Split(arg, "/")
			imagenameandtag := splitedArg[len(splitedArg)-1]

			if strings.Contains(imagenameandtag, ":") {
				image = strings.Split(imagenameandtag, ":")[0]
				if image == "" {
					logrus.Fatal("Image name cannot be nil")
				}

				tag = strings.Split(imagenameandtag, ":")[1]
				if tag == "" {
					tag = "latest"
				}
			} else {
				image = strings.Split(imagenameandtag, ":")[0]
				if image == "" {
					logrus.Fatal("Image name cannot be nil")
				}

				tag = "latest"
			}
		} else {
			logrus.Fatal("Please change imagename like username/image:[tag]")
		}

		var repository string

		//generate created
		var created = string(time.Now().Format("2006-01-02"))

		if pushPublicFlag {
			//push public image
			repository = "public"
		} else {
			//push private image
			repository = ConfigContent.Username
		}

		//generate arguments
		args0 := "curl"
		args1 := "-F"
		args2 := "username=" + ConfigContent.Username
		args3 := "-F"
		args4 := "password=" + ConfigContent.Password
		args5 := "-F"
		args6 := "repository=" + repository
		args7 := "-F"
		args8 := "tag=" + tag
		args9 := "-F"
		args10 := "imageID=" + imageID
		args11 := "-F"
		args12 := "created=" + created
		args13 := "-F"
		args14 := "size=" + size
		args15 := "-F"
		args16 := "imageType=" + imageType
		args17 := "-F"
		args18 := "owner=" + ConfigContent.Username
		args19 := "-F"
		args20 := image + "=@" + path
		args21 := ConfigContent.Apiserver + "/images/" + repository + "/" + image

		//execute curl to push image
		child := exec.Command(args0, args1, args2, args3, args4,
			args5, args6, args7, args8, args9, args10, args11,
			args12, args13, args14, args15, args16, args17,
			args18, args19, args20, args21)

		output, err := child.Output()
		if err != nil {
			logrus.Fatal(err)
		}

		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.SetUsageTemplate(pushUsage)
	pushCmd.Flags().StringVarP(&cfgFile, "configure-file", "f", "", "image's configure file (required)")
	pushCmd.Flags().StringVarP(&imageType, "image-type", "t", "docker", "Point out the type of image")
	pushCmd.MarkFlagRequired("image-type")
	// pushCmd.MarkFlagRequired("configure-file")
	pushCmd.Flags().BoolVarP(&pushPublicFlag, "public", "p", false, "Push a public image")
}
