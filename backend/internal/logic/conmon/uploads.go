package conmon

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"xiaozhu/backend/utils"
)

type UploadsIer interface {
	Save() (string, error)
	ValidateSize() error
	ValidateType() error
}

func SaveLogic(l UploadsIer) (string, error) {
	// 类型校验
	if err := l.ValidateType(); err != nil {
		return "", err
	}
	// 大小校验
	if err := l.ValidateSize(); err != nil {
		return "", err
	}
	return l.Save()

}

type UploadsLogic struct {
	ctx  context.Context
	File *multipart.FileHeader
}

func NewUploadsLogic(ctx context.Context, file *multipart.FileHeader) *UploadsLogic {
	return &UploadsLogic{
		ctx:  ctx,
		File: file,
	}
}

func (l *UploadsLogic) Save() (string, error) {
	ext := utils.GetFileExt(l.File.Filename)
	if ext == "" {
		return "", errors.New("文件异常，无扩展名")
	}

	// 临时存储文件路径
	rootPath := viper.GetString("oss.file")
	if err := utils.TidyDirectory(rootPath); err != nil {
		return "", err
	}

	// 先保存临时文件
	tmpFilePath := utils.NormalizePath(filepath.Join(rootPath, l.File.Filename))
	err := save(tmpFilePath, l.File)
	if err != nil {
		return "", err
	}

	// 删除临时文件
	defer func() {
		if _, err = os.Stat(tmpFilePath); err == nil {
			_ = os.Remove(tmpFilePath)
		}
	}()

	// 文件md5
	hashFilePath, err := utils.GetFileMd5(tmpFilePath)
	if err != nil {
		return "", err
	}

	// 最终保存文件名
	finalFilePath := utils.NormalizePath(filepath.Join(rootPath, hashFilePath+ext))

	// 检查是否存在
	_, err = os.Stat(finalFilePath)
	if err == nil { // 文件存在
		return finalFilePath, nil
	} else if !os.IsNotExist(err) { // 文件是非存在类型错误
		return "", err
	}

	// 重命名
	err = os.Rename(tmpFilePath, finalFilePath)
	if err != nil {
		return "", err
	}

	return finalFilePath, nil
}

func save(filepath string, file *multipart.FileHeader) error {
	tf, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer tf.Close()

	cf, err := file.Open()
	if err != nil {
		return err
	}
	defer cf.Close()
	_, err = io.Copy(tf, cf)
	if err != nil {
		return err
	}

	return nil
}

type ImageLogic struct {
	*UploadsLogic
}

func NewImage(ctx context.Context, file *multipart.FileHeader) *ImageLogic {
	return &ImageLogic{
		UploadsLogic: NewUploadsLogic(ctx, file),
	}
}

func (l *ImageLogic) ValidateType() error {
	ext := utils.GetFileExt(l.File.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".gif" {
		return errors.New("不支持的图片格式")
	}
	return nil
}

func (l *ImageLogic) ValidateSize() error {
	if l.File.Size > 10*1024*1024 {
		return errors.New("图片文件大小不能超过10m")
	}
	return nil
}

type VideoLogic struct {
	*UploadsLogic
}

func NewVideoLogic(ctx context.Context, file *multipart.FileHeader) *VideoLogic {
	return &VideoLogic{
		UploadsLogic: NewUploadsLogic(ctx, file),
	}
}

func (l *VideoLogic) ValidateType() error {
	ext := utils.GetFileExt(l.File.Filename)
	if ext != ".mp4" && ext != ".avi" && ext != ".mkv" {
		return errors.New("不支持的视频格式")
	}
	return nil
}

func (l *VideoLogic) ValidateSize() error {
	if l.File.Size > 100*1024*1024 {
		return errors.New("视频文件大小不能超过100m")
	}
	return nil
}

type FileLogic struct {
	*UploadsLogic
}

func NewFileLogic(ctx context.Context, file *multipart.FileHeader) *FileLogic {
	return &FileLogic{
		UploadsLogic: NewUploadsLogic(ctx, file),
	}
}

func (l *FileLogic) ValidateType() error {

	return nil
}

func (l *FileLogic) ValidateSize() error {
	if l.File.Size > 1024*1024*1024 {
		return errors.New("文件大小不能超过1G")
	}
	return nil
}
