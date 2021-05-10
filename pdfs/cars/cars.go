package cars

import (
	"fmt"
	"os"

	"github.com/eric-fleming/PDFmaker/data"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func PrintPdf(path string) {
	// initialize
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	// adding to pdf
	buildHeading(m)
	buildCarList(m)

	err := m.OutputFileAndClose(path + "Mitsubishi-2017-report.pdf")
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
}

func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("assets/images/Mitsubishi-2017.jpg", props.Rect{
					Center:  true,
					Percent: 50,
				})

				if err != nil {
					fmt.Println("Image file was not loaded 😱 - ", err)
				}

			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Mitsubishi 2017 Full Line", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getBlueColor(),
			})
		})
	})
}

func buildCarList(m pdf.Maroto) {
	tableHeadings := []string{"Make", "Model", "Type", "Price"}
	contents := data.CarList(8)
	lightBlueColor := getLightBlueColor()
	m.SetBackgroundColor(getBlueColor())

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Products", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 3, 2, 4},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 3, 2, 4},
		},
		Align:                consts.Left,
		HeaderContentSpace:   1,
		Line:                 false,
		AlternatedBackground: &lightBlueColor,
	})

}

func getBlueColor() color.Color {
	return color.Color{
		Red:   0,
		Green: 0,
		Blue:  90,
	}
}

func getLightBlueColor() color.Color {
	return color.Color{
		Red:   0,
		Green: 0,
		Blue:  200,
	}
}
