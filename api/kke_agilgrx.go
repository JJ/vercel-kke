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
		hito := strconv.Atoi(update.Message.CommandArguments())
		switch update.Message.Command() {
		case "kke":
			text = fmt.Sprintf( "→ Hito %s : %s\n🔗 https://jj.github.io/curso-tdd/temas/%s\n⚒ https://jj.github.io/curso-tdd/temas/%s#Actividad",
				hito,
				hitos[next].Title,
				hitos[next].URI,
				hitos[next].URI,
			)
		default:
			text = "Usa /kke <hito> para más información sobre el hito de ÁgilGRX correspondiente"
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
