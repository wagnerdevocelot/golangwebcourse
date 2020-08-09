=====================================================================================================

**Dir** tipo func

// Dir retorna todos, exceto o último elemento do caminho, normalmente o diretório do caminho.
// Depois de descartar o elemento final, Dir chama Clean no caminho e na trilha
// barras são removidas.
// Se o caminho estiver vazio, Dir retorna ".".
// Se o caminho consistir inteiramente em separadores, Dir retornará um único separador.
// O caminho retornado não termina em um separador, a menos que seja o diretório raiz.

func Dir(path string) string {
	vol := VolumeName(path)
	i := len(path) - 1
	for i >= len(vol) && !os.IsPathSeparator(path[i]) {
		i--
	}
	dir := Clean(path[len(vol) : i+1])
	if dir == "." && len(vol) > 2 {
		// deve ser UNC
		return vol
	}
	return vol + dir
}

=====================================================================================================
