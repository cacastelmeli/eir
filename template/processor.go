package template

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/cacastelmeli/eir/context"
	"github.com/cacastelmeli/eir/util"
)

type Processor struct {
	ctx              *context.Context
	wg               sync.WaitGroup
	contentTemplate  *contentTemplate
	filenameTemplate *filenameTemplate
}

func NewProcessor(ctx *context.Context) *Processor {
	return &Processor{
		ctx:              ctx,
		contentTemplate:  newContentTemplate(ctx),
		filenameTemplate: newFilenameTemplate(ctx),
	}
}

func collectTemplateFilenames(dir string) (filenames []string, err error) {
	files, err := os.ReadDir(dir)

	if err != nil {
		return
	}

	for _, file := range files {
		curDir := path.Join(dir, file.Name())

		if file.IsDir() {
			collectedFilenames, err := collectTemplateFilenames(curDir)

			if err != nil {
				return nil, err
			}

			filenames = append(filenames, collectedFilenames...)
		} else {
			filenames = append(filenames, curDir)
		}
	}

	return
}

func (processor *Processor) compileFilename(filename string) (string, error) {
	return processor.filenameTemplate.Compile(filename)
}

func (processor *Processor) compileContent(content string) (string, error) {
	return processor.contentTemplate.Compile(content)
}

func (processor *Processor) processFile(rootDir, targetPath, templateFilename string) error {
	templateFileContent, err := os.ReadFile(templateFilename)

	if err != nil {
		return err
	}

	resultFilename, err := processor.compileFilename(templateFilename)

	if err != nil {
		return err
	}

	resultFileContent, err := processor.compileContent(string(templateFileContent))

	if err != nil {
		return err
	}

	// Point to target path
	targetFilename := strings.Replace(resultFilename, rootDir, targetPath, 1)

	// Create directory
	if err = os.MkdirAll(path.Dir(targetFilename), os.ModePerm); err != nil {
		return nil
	}

	// Create file
	if file, err := os.Create(targetFilename); err != nil {
		return err
	} else {
		defer file.Close()

		// Write content to file
		_, err = file.WriteString(resultFileContent)
		return err
	}
}

func (processor *Processor) ProcessTemplate(templateName, targetPath string) error {
	rootDir := util.PathCacheTemplate(templateName)
	filenames, err := collectTemplateFilenames(rootDir)
	processErrors := []error{}

	if err != nil {
		return err
	}

	for _, filename := range filenames {
		processor.wg.Add(1)

		go func(filename string) {
			defer processor.wg.Done()

			if err := processor.processFile(rootDir, targetPath, filename); err != nil {
				processErrors = append(processErrors, err)
			}
		}(filename)
	}

	processor.wg.Wait()

	if errCount := len(processErrors); errCount > 0 {
		var nextMsg string
		firstErrMsg := processErrors[0].Error()

		if errCount > 1 {
			nextMsg = fmt.Sprintf("\n\nPlus other %v error(s)", errCount)
		}

		return errors.New(firstErrMsg + nextMsg)
	}

	return nil
}
