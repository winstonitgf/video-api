package services

import (
	"fmt"
	"mime/multipart"
	"time"
	"video-api/utils"

	"github.com/winstonitgf/video-process/cloudflare"
)

type CloudflareService struct {
	Error             error
	cloudflareSetting cloudflare.CloudflareSetting
}

func (c *CloudflareService) init() {
	c.cloudflareSetting.APIKey = utils.EnvConfig.Cloudflare.APIKey
	c.cloudflareSetting.Email = utils.EnvConfig.Cloudflare.Email
	c.cloudflareSetting.AccountID = utils.EnvConfig.Cloudflare.AccountID
	c.cloudflareSetting.APIDomain = utils.EnvConfig.Cloudflare.APIDomain
	c.cloudflareSetting.APIVersion = utils.EnvConfig.Cloudflare.APIVersion
	c.cloudflareSetting.Pem = utils.EnvConfig.Cloudflare.Pem
	c.cloudflareSetting.UtilDomain = utils.EnvConfig.Cloudflare.UtilDomain
	c.cloudflareSetting.KeyID = utils.EnvConfig.Cloudflare.KeyID
	c.cloudflareSetting.StreamDomain = utils.EnvConfig.Cloudflare.StreamDomain
}

func (c *CloudflareService) SearchVideo(videoName string) *cloudflare.VideoSearchResponse {

	cloudflareService := cloudflare.NewService(c.cloudflareSetting)
	videoSearchResponse, err := cloudflareService.Search(videoName)
	if err != nil {
		c.handleError(err)
		return nil
	}
	return &videoSearchResponse
}

func (c *CloudflareService) UploadVideo(fileHeader *multipart.FileHeader) *cloudflare.UploadReturnModel {

	meta := make(map[string]string)
	meta["name"] = fileHeader.Filename
	meta["requiresignedurls"] = "true"

	var uploadParameter cloudflare.UploadParameter
	uploadParameter.Filename = fileHeader.Filename
	uploadParameter.Fingerprint = fmt.Sprintf("%s-%d-%s", fileHeader.Filename, fileHeader.Size, time.Now())
	uploadParameter.Metadata = meta
	uploadParameter.Reader, _ = fileHeader.Open()
	uploadParameter.Size = fileHeader.Size

	cloudflareService := cloudflare.NewService(c.cloudflareSetting)
	uploadReturnModel, err := cloudflareService.Upload(uploadParameter)
	if err != nil {
		c.handleError(err)
		return nil
	}
	return &uploadReturnModel
}

func (c *CloudflareService) GetSignedURL(videoUID string) string {

	cloudflareService := cloudflare.NewService(c.cloudflareSetting)
	signedURL, err := cloudflareService.GetSignedURL(videoUID)
	if err != nil {
		c.handleError(err)
		return ""
	}
	return signedURL
}

func (c *CloudflareService) handleError(err error) {
	c.Error = err
}
