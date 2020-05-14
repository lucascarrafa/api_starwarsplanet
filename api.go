package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"context"
	"log"
	"io/ioutil"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
)

type Planeta struct {
	Id primitive.ObjectID `json:"_id, omitempty" bson:"_id,omitempty"`
    Nome string
	Clima string
	Terreno string
	Filmes int
}

type Results struct{
    Results []Info `json:"results"`
}

type Info struct
{
	Name string `json:"name"`
	Films []string `json:"films"`
}

func inserirPlaneta(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	client, err := mongo.NewClient(options.Client().ApplyURI("LINK_MONGOATLAS"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	universeDatabase := client.Database("universe")
	planetsCollection := universeDatabase.Collection("planets")

	var data Planeta
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	res, errGet := http.Get("https://swapi.dev/api/planets/?search="+data.Nome)
	if errGet != nil {
		fmt.Fprintf(w, "Ops! Inserir planeta válido")
	}

	b, _ := ioutil.ReadAll(res.Body)
	var m Results
	json.Unmarshal(b, &m)
	qtd := len(m.Results[0].Films)

	texto := Planeta{Nome: data.Nome, Clima: data.Clima, Terreno: data.Terreno, Filmes: qtd}
	resultInsert, _ := planetsCollection.InsertOne(ctx, texto)
	json.NewEncoder(w).Encode(resultInsert)
}

func listaPlanetas(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	client, err := mongo.NewClient(options.Client().ApplyURI("LINK_MONGOATLAS"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	universeDatabase := client.Database("universe")
	planetsCollection := universeDatabase.Collection("planets")

	var planets []Planeta

	cur, err := planetsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var planet Planeta
		err := cur.Decode(&planet)
		if err != nil {
			log.Fatal(err)
		}
		planets = append(planets, planet)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(planets)
}

func buscaNome(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	client, err := mongo.NewClient(options.Client().ApplyURI("LINK_MONGOATLAS"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	universeDatabase := client.Database("universe")
	planetsCollection := universeDatabase.Collection("planets")

	var planet Planeta
	vars := mux.Vars(r)
	filter := bson.M{"nome": vars["nome"] }
	planetsCollection.FindOne(ctx, filter).Decode(&planet)
	json.NewEncoder(w).Encode(planet)
}

func delPlaneta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	client, err := mongo.NewClient(options.Client().ApplyURI("LINK_MONGOATLAS"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	universeDatabase := client.Database("universe")
	planetsCollection := universeDatabase.Collection("planets")
	
	vars := mux.Vars(r)
	idPrimitive, errid := primitive.ObjectIDFromHex(vars["id"])
	if errid != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}
	filter := bson.M{"_id": idPrimitive}

	deleteResult, _ := planetsCollection.DeleteOne(ctx, filter)

	json.NewEncoder(w).Encode(deleteResult)
}
func buscaId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	client, err := mongo.NewClient(options.Client().ApplyURI("LINK_MONGOATLAS"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	universeDatabase := client.Database("universe")
	planetsCollection := universeDatabase.Collection("planets")
	vars := mux.Vars(r)
	idPrimitive, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}
	var planet Planeta
	filter := bson.M{"_id": idPrimitive }

	planetsCollection.FindOne(ctx, filter).Decode(&planet)
	json.NewEncoder(w).Encode(planet)
}

func status(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "API OK")
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", status).Methods("GET")
	r.HandleFunc("/add", inserirPlaneta).Methods("POST")
	r.HandleFunc("/lista", listaPlanetas).Methods("GET")
	r.HandleFunc("/busca/{nome}", buscaNome).Methods("GET")
	r.HandleFunc("/del/{id}", delPlaneta).Methods("DELETE")
	r.HandleFunc("/buscaID/{id}", buscaId).Methods("GET")
	http.ListenAndServe(":3333", r)
}

