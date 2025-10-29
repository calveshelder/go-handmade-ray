package main

import (
	"fmt"
)

func main() {

	// Image

	image_width := 256
	image_height := 256

	// Render
	
	fmt.Printf("P3\n%d %d\n255\n", image_width, image_height)

	for j := 0; j < image_height; j++ {
		for i := 0; i < image_width; i++ {
			var r float64 = float64(i) / float64(image_width - 1)
			var g float64 = float64(j) / float64(image_height - 1)
			b := 0.0;

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Printf("%d %d %d\n",ir, ig, ib)
		}
	}
	
}
