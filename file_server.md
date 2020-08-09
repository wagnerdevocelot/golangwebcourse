================================================================================
FUNÇÃO

**FileServer** tipo função < **FileSystem**

// FileServer retorna um manipulador que atende solicitações HTTP
// com o conteúdo do sistema de arquivos enraizado na raiz.
//
// Para usar a implementação do sistema de arquivos do sistema operacional,
// use http.Dir:
//
// http.Handle ("/", http.FileServer (http.Dir ("/ tmp")))
//
// Como um caso especial, o servidor de arquivos retornado redireciona qualquer solicitação
// terminando em "/index.html" para o mesmo caminho, sem o final
// "index.html".

func FileServer(root FileSystem) Handler

================================================================================

ARGUMENTO

**FileSystem** tipo interface < **File** tipo interface

// Um ​​FileSystem implementa acesso a uma coleção de arquivos nomeados.
// Os elementos em um caminho de arquivo são separados por barra ('/', U + 002F)
// caracteres, independentemente da convenção do sistema operacional do host.

type FileSystem interface {
	Open(name string) (File, error)
}

// Um ​​arquivo é retornado por um método Open de FileSystem e pode ser
// servido pela implementação FileServer.
// Os métodos devem se comportar da mesma forma que em um * os.File.

type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}

================================================================================

RETORNO

**fileHandler** tipo struct

type fileHandler struct {
	root FileSystem
}

**FileSystem** tipo custom

// Um ​​Dir implementa FileSystem usando o sistema de arquivos nativo restrito a um
// árvore de diretório específica.
//
// Enquanto o método FileSystem.Open leva caminhos separados por '/', uma string de Dir
// o valor é um nome de arquivo no sistema de arquivos nativo, não um URL, por isso é separado
// por filepath.Separator, que não é necessariamente '/'.
//
// Observe que Dir permitirá acesso a arquivos e diretórios começando com um
// ponto final, o que poderia expor diretórios confidenciais como um diretório .git ou
// arquivos sensíveis como .htpasswd. Para excluir arquivos com um ponto inicial,
// remova os arquivos / diretórios do servidor ou crie um FileSystem personalizado
// implementação.
//
// Um ​​Dir vazio é tratado como ".".

================================================================================

