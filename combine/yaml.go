package combine

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// yaml 合并
type YamlCombine struct {
	yaml map[string]interface{}
}

// 合并多个yaml文件
// files: 需要合并的yaml文件
func (c *YamlCombine) CombineFile(files ...string) error {
	if c.yaml == nil {
		c.yaml = make(map[string]interface{})
	}
	for _, file := range files {
		b, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		ym := make(map[string]interface{})
		err = yaml.Unmarshal(b, &ym)
		if err != nil {
			return err
		}
		//与all进行合并
		c.yaml = merge(c.yaml, ym)
	}

	return nil
}

// 合并指定目录下的yaml
func (c *YamlCombine) CombineDir(dir string) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".yaml") {
			if err := c.CombineFile(path); err != nil {
				return err
			}
		}
		return nil
	})
}

// merge 合并两个 map[string]interface{} 类型的映射
func merge(dst, src map[string]interface{}) map[string]interface{} {
	for key, value := range src {
		if v, ok := dst[key]; ok {
			switch v := v.(type) {
			case map[string]interface{}:
				if m, ok := value.(map[string]interface{}); ok {
					dst[key] = merge(v, m)
				} else {
					dst[key] = value
				}
			default:
				dst[key] = value
			}
		} else {
			dst[key] = value
		}
	}
	return dst
}

// 保存配置到文件
func (c *YamlCombine) Save(fileName string) error {
	bs, err := c.Bytes()
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, bs, 0644)
}

// 获取yaml文件内容
func (c *YamlCombine) Bytes() ([]byte, error) {
	if c.yaml == nil {
		return nil, errors.New("yaml is nil")
	}
	bs, err := yaml.Marshal(c.yaml)
	if err != nil {
		return nil, err
	}
	return bs, err
}

// 反序列化yaml到指定结构体
func (c *YamlCombine) Unmarshal(cy interface{}) error {
	bs, err := c.Bytes()
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(bs, cy); err != nil {
		return err
	}
	return nil
}
