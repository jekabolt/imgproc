package server

import (
	"bytes"
	"container/heap"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	quads "github.com/jekabolt/go-quads"
)

func (pr *ProcRouter) handleDotImage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("handleDotImage:ioutil.ReadAll: %v", err.Error())
		return
	}
	img, err := jpeg.Decode(bytes.NewReader(body))
	if err != nil {
		logger.Errorf("handleDotImage:jpeg.Decode: %v", err.Error())
		return
	}
	filename := strconv.Itoa(int(time.Now().Unix())) + ".jpg"
	out, err := os.Create("./out/" + filename)
	if err != nil {
		logger.Errorf("handleDotImage:os.Create: %v", err.Error())
		return
	}
	err = jpeg.Encode(out, img, &jpeg.Options{
		Quality: 100,
	})
	if err != nil {
		logger.Errorf("handleDotImage:jpeg.Encode: %v", err.Error())
		return
	}
	logger.Debugf("Image successfully saved [%v]", filename)

	img, err = cropImage(img, pr.ImageDimension)
	if err != nil {
		logger.Errorf("handleDotImage:cropImage: %v", err.Error())
		return
	}
	imgNRGBA, err := quads.ToNRGBA(img)
	if err != nil {
		logger.Errorf("handleDotImage:quads.DecodeImage: %v", err.Error())
		return
	}
	imgNRGBA = quads.NrgbaToGray(imgNRGBA)

	headNode, err := quads.InitializeFromFile(imgNRGBA)
	if err != nil {
		logger.Errorf("handleDotImage:quads.InitializeFromFile: %v", err.Error())
		return
	}

	mh := make(quads.MinHeap, 1)
	mh[0] = headNode
	heap.Init(&mh)
	img, err = quads.IterateV2(&mh, headNode, 1000, "0,0,0,255")
	if err != nil {
		logger.Errorf("handleDotImage:quads.IterateV2: %v", err.Error())
		return
	}

	//TODO: save quad iamge
	// filename = strconv.Itoa(int(time.Now().Unix())) + "Q.jpg"
	// out, err = os.Create("./out/" + filename)
	// if err != nil {
	// 	logger.Errorf("handleDotImage:os.Create: %v", err.Error())
	// 	return
	// }
	// jpeg.Encode(out, imgs, &jpeg.Options{
	// 	Quality: 100,
	// })

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}

}
