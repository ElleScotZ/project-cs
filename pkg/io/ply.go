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
	)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var (
		mesh       core.Mesh
		endHeader  bool = false
		lineNumber int
		scanner    = bufio.NewScanner(file)
	)

	// Reading in the header
	for !endHeader {
		scanner.Scan()
		line := scanner.Text()
		lineNumber++

		switch {
		case strings.Contains(line, vertex):
			vertexNumber, _ = strconv.Atoi(line[len(vertex)+1:])
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

			vertex := algebra.Vector3D{Position: [3]float64{posX, posY, posZ}}

			mesh.Vertices = append(mesh.Vertices, &vertex)

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

			mesh.Faces = append(mesh.Faces, &face)
		}
	}

	return &mesh, nil
}

// ExportPLY method writes out PLY files with fileName.
func ExportPLY(m *core.Mesh, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	// Every PLY file header
	_, err = file.WriteString("ply\nformat ascii 1.0\ncomment omitted\n")
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

	_, err = file.WriteString("element face " + strconv.Itoa(len(m.Faces)) + "\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("property list uchar int vertex_indices\nend_header\n")
	if err != nil {
		return err
	}

	for _, v := range m.Vertices {
		X := strconv.FormatFloat(v.Position[0], 'f', -1, 32)
		Y := strconv.FormatFloat(v.Position[1], 'f', -1, 32)
		Z := strconv.FormatFloat(v.Position[2], 'f', -1, 32)

		_, err := file.WriteString(X + " " + Y + " " + Z)
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
