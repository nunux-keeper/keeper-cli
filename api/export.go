package api

import (
	"errors"
	"io"
	"os"

	"gopkg.in/cheggaaa/pb.v1"
)

// DownloadExport Download export file
func (k *Client) DownloadExport(out *os.File) error {
	res, err := k.Get("/v2/exports", nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return errors.New(res.Status)
	}
	bar := pb.New(int(res.ContentLength)).SetUnits(pb.U_BYTES)
	bar.Start()

	// create proxy reader
	reader := bar.NewProxyReader(res.Body)
	_, err = io.Copy(out, reader)
	if err != nil {
		return err
	}
	return nil
}

// ScheduleExport Schedule a new export
func (k *Client) ScheduleExport() error {
	res, err := k.Post("/v2/exports", nil, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return errors.New(res.Status)
	}
	return nil
}

// GetExportStatus Get export status
func (k *Client) GetExportStatus(out io.Writer) error {
	res, err := k.Get("/v2/exports/status", nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(out, res.Body)
		return errors.New(res.Status)
	}
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}
