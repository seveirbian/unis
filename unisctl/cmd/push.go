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

		//generate created
		var created = string(time.Now().Format("2006-01-02"))
		i := string(time.Now().Format("2006-01-02T15:04:05Z"))

		//generate imageID
		if content, err := ioutil.ReadFile(path); err != nil {
			logrus.Fatal(err)
		} else {
			temp := sha256.Sum256([]byte(string(content) + i))
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

		if pushPublicFlag {
			//push public image
			repository = "public"
		} else {
			//push private image
			repository = ConfigContent.Username
		}

		//generate arguments
		arg0 := "curl"
		arg1 := "-F"
		arg2 := "username=" + ConfigContent.Username
		arg3 := "-F"
		arg4 := "password=" + ConfigContent.Password
		arg5 := "-F"
		arg6 := "repository=" + repository
		arg7 := "-F"
		arg8 := "tag=" + tag
		arg9 := "-F"
		arg10 := "imageID=" + imageID
		arg11 := "-F"
		arg12 := "created=" + created
		arg13 := "-F"
		arg14 := "size=" + size
		arg15 := "-F"
		arg16 := "imageType=" + imageType
		arg17 := "-F"
		arg18 := "owner=" + ConfigContent.Username
		arg19 := "-F"
		arg20 := image + "=@" + path
		arg21 := ConfigContent.Apiserver + "/images/push/" + repository + "/" + image

		//execute curl to push image
		child := exec.Command(arg0, arg1, arg2, arg3, arg4,
			arg5, arg6, arg7, arg8, arg9, arg10, arg11,
			arg12, arg13, arg14, arg15, arg16, arg17,
			arg18, arg19, arg20, arg21)

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
