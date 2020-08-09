================================================================================

FUNÇÃO

**listenAndServe** função

// ListenAndServe escuta no endereço de rede TCP addr e então chama 
// serve com o handler para lidar com solicitações em conexões de entrada. 
// As conexões aceitas são configuradas para habilitar keep-alives TCP. 
// 
// O handler é normalmente nil, caso em que o DefaultServeMux é usado. 
// 
// ListenAndServe sempre retorna um erro não nulo. 

func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

================================================================================

RETORNO

**Server** tipo struct

// Um ​​servidor define parâmetros para executar um servidor HTTP.
// O valor zero para Servidor é uma configuração válida.


================================================================================

ARGUMENTO 1

**addr** tipo string

// Addr especifica opcionalmente o endereço TCP para o servidor escutar,
// no formato "host: porta". Se estiver vazio, ": http" (porta 80) será usado.
// Os nomes dos serviços são definidos na RFC 6335 e atribuídos pela IANA.
// Consulte net.Dial para obter detalhes sobre o formato do endereço.

================================================================================

REFERÊNCIA

**servemux** tipo struct

ServeMux é um multiplexador de solicitação HTTP. Ele compara a URL de cada solicitação recebida com uma lista de padrões registrados e chama o manipulador para o padrão que mais se aproxima da URL.

Os padrões nomeiam caminhos fixos, com raiz, como "/favicon.ico", ou subárvores com raiz, como "/ imagens /" (observe a barra à direita). Os padrões mais longos têm precedência sobre os mais curtos, de modo que se houver manipuladores registrados para "/ images /" e "/ images / thumbnails /", o último manipulador será chamado para caminhos que começam com "/ images / thumbnails /" e o anterior receberá solicitações para quaisquer outros caminhos na subárvore "/ images /".

Observe que, uma vez que um padrão que termina em uma barra nomeia uma subárvore com raiz, o padrão "/" corresponde a todos os caminhos não correspondidos por outros padrões registrados, não apenas a URL com Caminho == "/".

Se uma subárvore foi registrada e uma solicitação é recebida nomeando a raiz da subárvore sem sua barra final, o ServeMux redireciona essa solicitação para a raiz da subárvore (adicionando a barra final). Esse comportamento pode ser substituído por um registro separado para o caminho sem a barra final. Por exemplo, registrar "/ images /" faz com que o ServeMux redirecione uma solicitação de "/ images" para "/ images /", a menos que "/ images" tenha sido registrado separadamente.

Os padrões podem opcionalmente começar com um nome de host, restringindo correspondências a URLs apenas nesse host. Os padrões específicos do host têm precedência sobre os padrões gerais, de modo que um manipulador pode se registrar para os dois padrões "/ codesearch" e "codesearch.google.com/" sem também assumir as solicitações para " http://www.google.com/ "

O ServeMux também se encarrega de limpar o caminho da solicitação de URL e o cabeçalho do Host, removendo o número da porta e redirecionando qualquer solicitação que o contenha. ou .. elementos ou barras repetidas para um URL equivalente e mais limpo.

type ServeMux struct {
    // contém campos filtrados ou não exportados 
}

**DefaultServeMux**

DefaultServeMux é o ServeMux padrão usado pelo Serve.

================================================================================

ARGUMENTO 2

**Handler** tipo interface

// Um ​​Handler responde a uma solicitação HTTP.
//
// ServeHTTP deve escrever cabeçalhos de resposta e dados no ResponseWriter
// e depois retorna. Retornar sinaliza que a solicitação está concluída; isto
// não é válido usar o ResponseWriter ou ler do
// Request.Body após ou simultaneamente com a conclusão do
// Atende chamada de HTTP.
//
// Dependendo do software cliente HTTP, versão do protocolo HTTP e
// quaisquer intermediários entre o cliente e o servidor Go, ele não pode
// seja possível ler o Request.Body depois de escrever no
// ResponseWriter. Manipuladores cautelosos devem ler o Request.Body
// primeiro, e depois responda.
//
// Exceto para ler o corpo, os manipuladores não devem modificar o
// Solicitação fornecida.
//
// Se o ServeHTTP entrar em pânico, o servidor (o chamador do ServeHTTP) assumirá
// que o efeito do pânico foi isolado à solicitação ativa.
// Ele recupera o pânico, registra um rastreamento de pilha no log de erros do servidor,
// e fecha a conexão de rede ou envia um HTTP / 2
// RST_STREAM, dependendo do protocolo HTTP. Abortar um manipulador para
// o cliente vê uma resposta interrompida, mas o servidor não registra
// um erro, entre em pânico com o valor ErrAbortHandler.

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

================================================================================