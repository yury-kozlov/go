# Examples of public func types defined in GO sources:
  type Compressor func(w io.Writer) (io.WriteCloser, error)
  type Decompressor func(r io.Reader) io.ReadCloser
  type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
  type FuncType func(x, y int, s string) (b *B, err error)
  type ErrorHandler func(err error)
  type PragmaHandler func(pos Pos, blank bool, text string, current Pragma) Pragma
  type ProgAlloc func() *Prog
  type LookupFn func(name string, version int) *sym.Symbol
  type FuncReturningInt func() int
  type PostProcessor func(input io.Reader, output io.Writer, ui plugin.UI) error
  type TagMatch func(s *Sample) bool
  type SymLookup func(uint64) (string, uint64)
  type VersionFixer func(path, version string) (string, error)
  type Hash func(files []string, open func(string) (io.ReadCloser, error)) (string, error)
  type HashReaderFunc func([]int64) ([]Hash, error)
  type ApplyFunc func(*Cursor) bool
  type CancelFunc func()
  type Func func() interface{}
  type Filter func(string) bool
  type FieldFilter func(name string, value reflect.Value) bool
  type Importer func(imports map[string]*Object, path string) (pkg *Object, err error)
  type Lookup func(path string) (io.ReadCloser, error)
  type Qualifier func(*Package) string
  type Demangler func(name []string) (map[string]string, error)
  type RoundTripperFunc func(*http.Request) (*http.Response, error)
  type HandlerFunc func(ResponseWriter, *Request)
  type AfterFilter func(*Status) error
  type WalkFunc func(path string, info os.FileInfo, err error) error
  type FuncDDD func(...interface{}) error
  type MyFunc func()
  type By func(p1, p2 *Planet) bool
  type BuilderContinuation func(child *Builder)
  type Option func(*options)
  type MakePipe func() (c1, c2 net.Conn, stop func(), err error)
