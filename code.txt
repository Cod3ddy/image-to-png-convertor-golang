package main

import (
    "fmt"
    "image"
    "image/color"
    _ "image/jpeg" // Import for JPEG format
    "os"
)

func main() {
    // Open the image file
    file, err := os.Open("input.jpg") // Change "input.jpg" to your image file
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Decode the image
    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Get the bounds of the image
    bounds := img.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y

    // Extract colors
    var colors []color.Color
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            pixelColor := img.At(x, y)
            colors = append(colors, pixelColor)
        }
    }

    // Create color palette (Here you can implement a clustering algorithm or simply choose dominant colors)
    // For simplicity, let's just pick the first 10 colors
    var palette []color.Color
    for i := 0; i < 10 && i < len(colors); i++ {
        palette = append(palette, colors[i])
    }

    // Generate mosaic effect
    // Map each pixel of the original image to the nearest color in the palette
    mosaic := image.NewRGBA(bounds)
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            originalColor := img.At(x, y)
            closestColor := findClosestColor(originalColor, palette)
            mosaic.Set(x, y, closestColor)
        }
    }

    // Save the mosaic image to file
    mosaicFile, err := os.Create("mosaic.jpg") // Change "mosaic.jpg" to your desired output file
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer mosaicFile.Close()
    jpeg.Encode(mosaicFile, mosaic, nil)

    fmt.Println("Mosaic effect generated successfully!")
}

// Function to find the closest color in the palette
func findClosestColor(target color.Color, palette []color.Color) color.Color {
    tr, tg, tb, _ := target.RGBA()
    minDist := uint32(1<<32 - 1)
    var closest color.Color

    for _, c := range palette {
        cr, cg, cb, _ := c.RGBA()
        dr := tr - cr
        dg := tg - cg
        db := tb - cb
        dist := dr*dr + dg*dg + db*db
        if dist < minDist {
            minDist = dist
            closest = c
        }
    }

    return closest
}
