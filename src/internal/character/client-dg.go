package character

import ("io"
	"log"
	"fmt"
	"net/http"
	"context"
	"time"
	"net/url"
    "encoding/json"
)



// dragonBallHTTPClient es el struct que hablará con la API externa.
type dragonBallHTTPClient struct {
	baseURL string
	http    *http.Client
}

func (d dragonBallHTTPClient) FetchByName(ctx context.Context, name string) (*Character, error) {
    // simula demora y "no encontrado"
     urlApi:= "https://dragonball-api.com/api/characters"
	// Para evitar errores de caracteres no url en la concatenacion normal 
	u, err := url.Parse(urlApi)
	if err != nil {
		panic(err)
	}
	// 2. Añades o modificas los parámetros
	q := u.Query()       
	q.Set("name", name)  
	u.RawQuery = q.Encode()

	// 3. URL final
	fullURL := u.String()
	fmt.Println(fullURL)


    client := &http.Client{Timeout: 10 * time.Second}

    // 2. Hacer GET
	resp, err := client.Get(fullURL)
	if err != nil {
		log.Fatalf("request: %v", err)
	}
    // en caso de fallas cerrar conexion (buen dato!! )
	defer resp.Body.Close()

	// 3. Validar status HTTP
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API respondió %s", resp.Status)
	}

    body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("leer body: %v", err)
	}

	// Imprimir como texto
	fmt.Println(string(body))	// 4. Decodificar JSON
	var chars []Character
	if err := json.Unmarshal(body, &chars); err != nil {
		log.Fatalf("unmarshal: %v", err)
	}

	// 5. Mostrar resultado
	for _, ch := range chars {
		fmt.Printf("%d - %s (%s)  Ki: %s\n", ch.ID, ch.Name, ch.Race, ch.Ki)
	}
	if len(chars) == 0 {
    	return nil, ErrNotFound  
	}
    return &chars[0],nil
   
}
// NewDragonBallHTTPClient devuelve un adaptador que satisface ExternalClient.
// ▸ Por ahora no implementa lógica; devuelve nil o un stub.
func NewDragonBallHTTPClient() ExternalClient {
	// TODO: construir y devolver un *dragonBallHTTPClient real
	return dragonBallHTTPClient{}
}
