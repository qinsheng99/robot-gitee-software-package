package pullrequestimpl

import (
	"bytes"
	"html/template"
)

type sigInfoTplData struct {
	PkgName       string
	ImporterEmail string
	Importer      string
}

type repoYamlTplData struct {
	PkgName     string
	PkgDesc     string
	BranchName  string
	ProtectType string
	PublicType  string
}

type prBodyTplData struct {
	PkgName string
	PkgLink string
}

func newtemplateImpl(cfg *Template) (templateImpl, error) {
	r := templateImpl{}

	// pr body
	tmpl, err := template.ParseFiles(cfg.PRBodyTpl)
	if err != nil {
		return r, err
	}
	r.prBodyTpl = tmpl

	// repo yaml
	tmpl, err = template.ParseFiles(cfg.RepoYamlTpl)
	if err != nil {
		return r, err
	}
	r.repoYamlTpl = tmpl

	// sig info
	tmpl, err = template.ParseFiles(cfg.SigInfoTpl)
	if err != nil {
		return r, err
	}
	r.sigInfoTpl = tmpl

	return r, nil
}

type templateImpl struct {
	prBodyTpl   *template.Template
	sigInfoTpl  *template.Template
	repoYamlTpl *template.Template
}

func (impl *templateImpl) genPRBody(data *prBodyTplData) (string, error) {
	return impl.gen(impl.prBodyTpl, data)
}

func (impl *templateImpl) genSigInfo(data *sigInfoTplData) (string, error) {
	return impl.gen(impl.sigInfoTpl, data)
}

func (impl *templateImpl) genRepoYaml(data *repoYamlTplData) (string, error) {
	return impl.gen(impl.repoYamlTpl, data)
}

func (impl *templateImpl) gen(tpl *template.Template, data interface{}) (string, error) {
	buf := new(bytes.Buffer)

	if err := tpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
