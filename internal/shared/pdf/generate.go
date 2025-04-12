package pdf

import (
	"strings"

	"github.com/signintech/gopdf"
)

func GeneratePDF(text, outputPath string) error {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	pdf.SetX(40)
	pdf.SetY(50)
	err := pdf.AddTTFFont("roboto", "./uploads/Roboto-Regular.ttf")
	if err != nil {
		return err
	}
	err = pdf.SetFont("roboto", "", 14)
	if err != nil {
		return err
	}

	paragraphs := strings.Split(text, "\n")
	for _, para := range paragraphs {
		lines := wrapText(pdf, para, 500) // 120mm width
		for _, line := range lines {
			pdf.Cell(nil, line)
			pdf.Br(16)
		}
		// extra space between paragraphs
		pdf.Br(8)
	}
	return pdf.WritePdf(outputPath)
}

func wrapText(pdf *gopdf.GoPdf, text string, maxWidth float64) []string {
	var lines []string
	words := strings.Fields(text)
	if len(words) == 0 {
		return lines
	}

	line := ""
	for _, word := range words {
		testLine := line
		if testLine != "" {
			testLine += " "
		}
		testLine += word

		width, _ := pdf.MeasureTextWidth(testLine)
		if width > maxWidth && line != "" {
			lines = append(lines, line)
			line = word
		} else {
			line = testLine
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}
