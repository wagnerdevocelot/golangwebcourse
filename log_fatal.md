
================================================================================
**Fatal** função
é equivalente a Print() seguido por uma chamada para os.Exit (1).

func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

================================================================================

**Output** função

// Output grava a saída para um evento de registro. A string s contém
// o texto a ser impresso após o prefixo especificado pelas bandeiras do
// Logger. Uma nova linha é acrescentada se o último caractere de s não for
// já é uma nova linha. Calldepth é a contagem do número de
// quadros a pular ao calcular o nome do arquivo e o número da linha
// se Llongfile ou Lshortfile estiver definido; um valor de 1 imprimirá os detalhes
// para o chamador de Output.

func Output(calldepth int, s string) error {
	return std.Output(calldepth+1, s) // +1 para este quadro.
}

================================================================================

**std** tipo var

// std é uma variavel de em package level no pacote de log
var std = New(os.Stderr, "", LstdFlags)

================================================================================

**Stderr** tipo var

// Stdin, Stdout e Stderr são arquivos abertos que apontam para a entrada padrão, 
// saída padrão e descritores de arquivo de erro padrão. 
// 
// Observe que o tempo de execução Go grava no erro padrão para pânicos e travamentos; 
// fechar o Stderr pode fazer com que essas mensagens vão para outro lugar, talvez 
// para um arquivo aberto mais tarde.
var (
    Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)

================================================================================

 **Exit** tipo func

faz com que o programa atual saia com o código de status fornecido.
// Convencionalmente, o código zero indica sucesso, diferente de zero um erro.
// O programa termina imediatamente; funções adiadas não são executadas.
//
// Para portabilidade, o código de status deve estar no intervalo [0, 125].

func Exit(code int) {
	if code == 0 {
		// Dê ao detector de corrida uma chance de falhar no programa.
		// Os programas Racy não têm o direito de terminar com sucesso.
		runtime_beforeExit()
	}
	syscall.Exit(code)
}

================================================================================

**syscall.exit** tipo func
//go:linkname syscall_Exit syscall.Exit
//go:nosplit
func syscall_Exit(code int) {
    exit(int32(code))
}

================================================================================

