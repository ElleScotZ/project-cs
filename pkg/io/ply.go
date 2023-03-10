package io

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/ElleScotZ/project-cs/internal/algebra"
	"github.com/ElleScotZ/project-cs/internal/core"
)

const (
	// outFolder is the string for the out folder.
	outFolder = "out/"
)

// ImportPLY method reads in PLY files with fileName.
// fileName requirements: ending .ply.
// About PLY files: http://paulbourke.net/dataformats/ply/
// It must be an ASCII PLY file.
func ImportPLY(fileName string) (*core.Mesh, error) {
	// Important PLY file header data
	var vertexNumber int

	const (
		header = "end_header"
		vertex = "element vertex"
		normal = "property float nx"
		colour = "property uchar red"
	)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var (
		mesh           core.Mesh
		endHeader      bool = false
		lineNumber     int
		colorIncluded  bool
		normalIncluded bool
		scanner        = bufio.NewScanner(file)
	)

	// Reading in the header
	for !endHeader {
		scanner.Scan()
		line := scanner.Text()
		lineNumber++

		switch {
		case strings.Contains(line, vertex):
			vertexNumber, _ = strconv.Atoi(line[len(vertex)+1:])
		case strings.Contains(line, normal):
			normalIncluded = true
		case strings.Contains(line, colour):
			colorIncluded = true
		case strings.Contains(line, header):
			endHeader = true
			lineNumber = 0
		}
	}

	// Reading in the data
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		lineNumber++

		switch {
		case lineNumber <= vertexNumber:
			line = strings.TrimSpace(line)
			coordinates := strings.Split(line, " ")

			size := len(coordinates)
			if (size % 3) != 0 {
				return nil, errors.New("invalid PLY file")
			}

			posX, err := strconv.ParseFloat(coordinates[0], 64)
			if err != nil {
				return &mesh, err
			}

			posY, err := strconv.ParseFloat(coordinates[1], 64)
			if err != nil {
				return &mesh, err
			}

			posZ, err := strconv.ParseFloat(coordinates[2], 64)
			if err != nil {
				return &mesh, err
			}

			vertex := core.Vertex{Position: algebra.Vector3D{Coordinates: [3]float64{posX, posY, posZ}}}

			if normalIncluded {
				if len(coordinates) < 6 {
					return nil, errors.New("invalid PLY file")
				}

				normX, err := strconv.ParseFloat(coordinates[3], 64)
				if err != nil {
					return &mesh, err
				}

				normY, err := strconv.ParseFloat(coordinates[4], 64)
				if err != nil {
					return &mesh, err
				}

				normZ, err := strconv.ParseFloat(coordinates[5], 64)
				if err != nil {
					return &mesh, err
				}

				vertex.Normal = algebra.Vector3D{Coordinates: [3]float64{normX, normY, normZ}}
			}

			if colorIncluded {
				if len(coordinates) == 3 {
					return nil, errors.New("invalid PLY file")
				}

				colR, err7 := strconv.ParseFloat(coordinates[size-3], 64)
				colG, err8 := strconv.ParseFloat(coordinates[size-2], 64)
				colB, err9 := strconv.ParseFloat(coordinates[size-1], 64)

				switch {
				case err7 != nil:
					return &mesh, err7
				case err8 != nil:
					return &mesh, err8
				case err9 != nil:
					return &mesh, err9
				}

				vertex.Color = algebra.Vector3D{Coordinates: [3]float64{colR, colG, colB}}
			}

			mesh.Vertices = append(mesh.Vertices, vertex)

		case lineNumber > vertexNumber:
			faceString := strings.Split(line, " ")

			vertex1, err1 := strconv.Atoi(faceString[1])
			vertex2, err2 := strconv.Atoi(faceString[2])
			vertex3, err3 := strconv.Atoi(faceString[3])

			if err1 != nil || err2 != nil || err3 != nil {
				return nil, errors.New("Invalid PLY file.")
			}

			face := core.Face{
				Vertices: [3]int{vertex1, vertex2, vertex3},
			}

			mesh.Faces = append(mesh.Faces, face)
		}
	}

	return &mesh, nil
}

// ExportPLY method writes out ASCII PLY files with fileName.
// Open file with Print 3D, Meshlab, Meshmixer.
func ExportPLY(m *core.Mesh, fileName string) error {
	err := os.MkdirAll("out", os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(outFolder + fileName + ".ply")
	if err != nil {
		return err
	}

	defer file.Close()

	// Every PLY file header
	_, err = file.WriteString("ply\nformat ascii 1.0\ncomment nothing at all\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("element vertex " + strconv.Itoa(len(m.Vertices)) + "\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("property float x\nproperty float y\nproperty float z\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("property float nx\nproperty float ny\nproperty float nz\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("property uchar red\nproperty uchar green\nproperty uchar blue\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("element face " + strconv.Itoa(len(m.Faces)) + "\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("property list uchar int vertex_indices\nend_header\n")
	if err != nil {
		return err
	}

	for _, v := range m.Vertices {
		X := strconv.FormatFloat(v.Position.Coordinates[0], 'f', -1, 32)
		Y := strconv.FormatFloat(v.Position.Coordinates[1], 'f', -1, 32)
		Z := strconv.FormatFloat(v.Position.Coordinates[2], 'f', -1, 32)

		_, err := file.WriteString(X + " " + Y + " " + Z)
		if err != nil {
			return err
		}

		nX := strconv.FormatFloat(v.Normal.Coordinates[0], 'f', -1, 32)
		nY := strconv.FormatFloat(v.Normal.Coordinates[1], 'f', -1, 32)
		nZ := strconv.FormatFloat(v.Normal.Coordinates[2], 'f', -1, 32)

		_, err = file.WriteString(" " + nX + " " + nY + " " + nZ)
		if err != nil {
			return err
		}

		cX := strconv.FormatFloat(v.Color.Coordinates[0], 'f', -1, 32)
		cY := strconv.FormatFloat(v.Color.Coordinates[1], 'f', -1, 32)
		cZ := strconv.FormatFloat(v.Color.Coordinates[2], 'f', -1, 32)

		_, err = file.WriteString(" " + cX + " " + cY + " " + cZ)
		if err != nil {
			return err
		}

		_, err = file.WriteString("\n")
		if err != nil {
			return err
		}
	}

	for _, f := range m.Faces {
		A := strconv.Itoa(f.Vertices[0])
		B := strconv.Itoa(f.Vertices[1])
		C := strconv.Itoa(f.Vertices[2])

		_, err = file.WriteString("3 " + A + " " + B + " " + C + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
