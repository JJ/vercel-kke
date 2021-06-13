package handler

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
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

var hitos = []Hito {
	Hito {
		URI: "git",
		Title: "Datos básicos y repo",
	},
	Hito {
		URI: "ddd",
		Title: "Idea/problema a resolver, «personas»",
	},
	Hito {
		URI: "aplicaciones",
		Title: "Épicas/Tipos",
	},
	Hito {
		URI: "servicios",
		Title: "Servicios en la nube",
	},
	Hito {
		URI: "ágil",
		Title: "Organización de un proyecto",
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
	Hito {
		URI: "CI",
		Title: "Sistemas de integración continua",
	},
	Hito {
		URI: "inversión",
		Title: "Inversión/inyección de dependencias, mocks",
	},
	Hito {
		URI: "cobertura",
		Title: "Tests de cobertura de caminos de código",
	},

}


func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	if err := json.Unmarshal(body,&update); err != nil {
		log.Fatal("Error en el update →", err)
	}
	if update.Message != nil &&  update.Message.IsCommand() && update.Message.Command() == "kke" {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		text := ""
		argument := update.Message.CommandArguments()
		log.Printf( "Argumento → [%s]", argument )
		hito, error := strconv.Atoi( argument )
		if error != nil {
			log.Printf("Argumento incorrecto → %s", argument )
			text = "El argumento no es correcto, usa /kke <número>"
		} else {
			switch update.Message.Command() {
			case "kke":
				if hito >=  len(hitos) {
					text = fmt.Sprintf( "No tenemos info sobre el hito %d", hito )
				} else {
					text = fmt.Sprintf( "→ Hito %d : %s\n🔗 https://jj.github.io/curso-tdd/temas/%s\n⚒ https://jj.github.io/curso-tdd/temas/%s#actividad",
					hito,
					hitos[hito].Title,
					hitos[hito].URI,
					hitos[hito].URI,
					)
				}
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
