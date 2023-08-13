package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	fakename "github.com/malikbenkirane/go-france"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var cmdGenerate = &cobra.Command{
	Use:   "generate",
	Short: "generate go files with france firstnames and lastnames",
	RunE: func(cmd *cobra.Command, args []string) error {
		var eg errgroup.Group
		dst := make(chan fakename.Nameset, 2)
		for _, s := range []struct {
			label string
			src   string
		}{
			{
				label: "lastnames",
				src:   flagLastnamesFile,
			},
			{
				label: "firstnames",
				src:   flagFirstnamesFile,
			},
		} {
			c := s
			eg.Go(func() error {
				f, err := os.Open(c.src)
				if err != nil {
					return fmt.Errorf("%s: os: open: %w", c.label, err)
				}
				defer f.Close()
				ns, err := fakename.NewSet(f, fakename.NamesetWithLabel(c.label)) // label will be overwritten with staticNnames
				if err != nil {
					return fmt.Errorf("%s: fakename: new set: %w", c.label, err)
				}
				dst <- ns
				return nil
			})
		}
		if err := eg.Wait(); err != nil {
			return err
		}
		fn, ln := func() (fn, ln fakename.Nameset) {
			ns := <-dst
			if ns.Label() == "firstnames" {
				fn = ns
				ln = <-dst
				return
			}
			ln = ns
			fn = <-dst
			return
		}()
		c := newTemplateCompiler(fn, ln)
		out, err := os.Create(flagGenOutput)
		if err != nil {
			return err
		}
		return c.compile(out)
	},
}

type tplc struct {
	fn, ln *tpls
	tpl    string
}

type tpls struct {
	s   fakename.Nameset
	cn  string
	cc  string
	cum int
}

func newTemplateCompiler(fn, ln fakename.Nameset) (c tplc) {
	c.fn, c.ln = &tpls{s: fn}, &tpls{s: ln}
	c.tpl = `package fakename

func staticNames(opts ...NamesetOption) (fn Nameset, ln Nameset) {
	conf := defaultNamesetConfig()
	for _, opt := range opts {
		conf = opt(conf)
	}
	conf.label = "firstnames"
	fn = nameset{
		conf:  conf,
		name:  []string{%s},
		count: []int{%s},
		cum:   %d,
	}
	conf.label = "lastnames"
	ln = nameset{
		conf:  conf,
		name:  []string{%s},
		count: []int{%s},
		cum:   %d,
	}
	return
}`
	return
}

func (c tplc) compile(w io.Writer) error {
	for _, s := range []*tpls{c.fn, c.ln} {
		n, c, cum := s.s.All()
		fmt.Println(s.s.Label(), ": count:", len(n))
		nl := make([]string, len(n))
		cl := make([]string, len(n)) // len(n) == len(c)
		for i := range n {
			nl[i] = fmt.Sprintf("%q", n[i])
			cl[i] = fmt.Sprintf("%d", c[i])
		}
		s.cn = strings.Join(nl, `, `)
		s.cc = strings.Join(cl, `, `)
		s.cum = cum
	}
	_, err := fmt.Fprintf(w, c.tpl, c.fn.cn, c.fn.cc, c.fn.cum, c.ln.cn, c.ln.cc, c.ln.cum)
	return err
}
