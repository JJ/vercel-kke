package handler

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"io/ioutil"
)

type Hito struct {
	URI  string
	Title string
	fecha time.Time
}

type Response struct {
	Msg string `json:"text"`
	ChatID int64 `json:"chat_id"`
	Method string `json:"method"`
}

func format_date(diff time.Duration) string {
	// Función para convertir una diferencia de fechas
	// a una cadena del tipo dias horas minutos segundos

    // Definimos constantes para partir la diferencia
    const Decisecond = 100 * time.Millisecond
    const Day = 24*time.Hour

    // Extraemos la cantidad de dias y las quitamos de la diferencia
    d := diff / Day
    diff = diff % Day
    // Extraemos la cantidad de horas
    h := diff / time.Hour
    diff = diff % time.Hour
    // Extraemos la cantidad de minutos
    m := diff / time.Minute
    diff = diff % time.Minute
    // Extraemos la cantidad de secundos
    s := diff / time.Second
    diff = diff % time.Second
    // Nos quedamos con las partes de segundo
    f := diff / Decisecond
    return  fmt.Sprintf("%dd %dh %dm %d.%ds", d, h, m, s, f)

}


var hitos = []Hito {
	Hito {
		URI: "git",
		Title: "Datos básicos y repo",
	},
	Hito {
		URI: "ágil",
		Title: "Idea/problema a resolver, «personas»",
		fecha: time.Date(2020, time.October, 6, 11, 30, 0, 0, time.UTC),
	},
	Hito {
		URI: "aplicaciones",
		Title: "Épicas",
	},
	Hito {
		URI: "servicios",
		Title: "Servicios en la nube",
	},
	Hito {
		URI: "diseño",
		Title: "Creando historias de usuario",
	},
	Hito {
		URI: "organizando",
		Title: "Planificación en Milestones",
	},
	Hito {
		URI: "a-programar",
		Title: "Diseño general de clases, excepciones, modularización",
	},
	Hito {
		URI: "gestores-tareas",
		Title: "Configuración como código: gestores de tareas",
	},
	Hito {
		URI: "hacia-tests-unitarios",
		Title: "Calidad en el código, linters",
	},
	Hito {
		URI: "tests-unitarios-organización",
		Title: "Bibliotecas de aserciones, setup",
	},
	Hito {
		URI: "tests-unitarios",
		Title: "Marcos de test",
	},

}


func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	if err := json.Unmarshal(body,&update); err != nil {
		log.Fatal("Error en el update →", err)
	}
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	if update.Message.IsCommand() {
		text := ""
		switch update.Message.Command() {
		case "kke":
			text = fmt.Sprintf( "→ Hito %s\n🔗 https://jj.github.io/IV/documentos/proyecto/%s\n📅 %s",
					hitos[next].Title,
					hitos[next].URI,
					hitos[next].fecha.String(),
				)
			default:
				text = "Usa /kke <hito> para más información sobre el hito de ÁgilGRX correspondiente"
			}
		}
		data := Response{ Msg: text,
			Method: "sendMessage",
			ChatID: update.Message.Chat.ID }

		msg, _ := json.Marshal( data )
		log.Printf("Response %s", string(msg))
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w,string(msg))
	}
}
