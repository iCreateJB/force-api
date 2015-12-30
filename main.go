package main

import (
	"encoding/json"
	"log"
  "os"
	"math/rand"
	"net/http"
	"time"
	"io/ioutil"

	"github.com/bmizerany/pat"
	// "github.com/jinzhu/gorm"
	// _ "github.com/lib/pq"
)

const appVersionStr = "1.0"

// var (
// 	db gorm.DB
// )

var morals = []string{
	"Great leaders inspire greatness in others.",
	"Belief is not a matter of choice, but of conviction.",
	"Easy is the path to wisdom for those not blinded by ego",
	"Easy is the path to wisdom for those not blinded by themselves",
	"A plan is only as good as those who see it through.",
	"The best confidence builder is experience.",
	"Trust in your friends, and they'll have reason to trust in you.",
	"You hold onto friends by keeping your heart a little softer than your head.",
	"Heroes are made by the times.",
	"Ignore your instincts at your peril.",
	"Most powerful is he who controls his own power.",
	"The winding path to peace is always a worthy one, regardless of how many turns it takes.",
	"Fail with honor rather than succeed by fraud.",
	"Greed and fear of loss are the roots that lead to the tree of evil.",
	"When surrounded by war, one must eventually choose a side.",
	"Arrogance diminishes wisdom.",
	"Truth enlightens the mind, but won't always bring happiness to your heart.",
	"Fear is a disease; hope is its only cure.",
	"A single chance is a galaxy of hope.",
	"It is a rough road that leads to the heights of greatness.",
	"The costs of war can never be truly accounted for.",
	"Compromise is a virtue to be cultivated, not a weakness to be despised.",
	"A secret shared is a trust formed.",
	"Easy is the path to wisdom for those not blinded by themselves",
	"A lesson learned is a lesson earned.",
	"Overconfidence is the most dangerous form of carelessness.",
	"The first step to correcting a mistake is patience.",
	"A true heart should never be doubted.",
	"Believe in yourself or no one else will.",
	"No gift is more precious than trust.",
	"Sometimes, accepting help is harder than offering it.",
	"Attachment is not compassion.",
	"For everything you gain, you lose something else.",
	"It is the quest for honor that makes one honorable.",
	"Easy isn't always simple.",
	"If you ignore the past, you jeopardize your future.",
	"Fear not for the future, weep not for the past.",
	"In war, truth is the first casualty.",
	"Searching for the truth is easy. Accepting the truth is hard.",
	"A wise leader knows when to follow.",
	"Courage makes heroes, but trust builds friendship.",
	"Choose what is right, not what is easy",
	"The most dangerous beast is the beast within.",
	"Who my father was matters less than my memory of him.",
	"Adversity is friendship's truest test.",
	"Revenge is a confession of pain.",
	"Brothers in arms are brothers for life.",
	"Fighting a war tests a soldier's skills, defending his home tests a soldier's heart.",
	"Where there's a will, there's a way.",
	"A child stolen is a lost hope.",
	"The challenge of hope is to overcome corruption.",
	"Those who enforce the law must obey the law.",
	"The future has many paths -- choose wisely.",
	"A failure in planning is a plan for failure.",
	"Love comes in all shapes and sizes.",
	"Fear is a great motivator.",
	"Truth can strike down the specter of fear",
	"The swiftest path to destruction is through vengeance.",
	"Evil is not born, it is taught.",
	"The path to evil may bring great power, but not loyalty.",
	"Balance is found in the one who faces his guilt.",
	"He who surrenders hope, surrenders life.",
	"He who seeks to control fate shall never find peace.",
	"Adaptation is the key to survival.",
	"Anything that can go wrong will.",
	"Without honor, victory is hollow.",
	"Without humility, courage is a dangerous game.",
	"A great student is what the teacher hopes to be.",
	"When destiny calls, the chosen have no choice.",
	"Only through fire is a strong sword forged.",
	"Crowns are inherited, kingdoms are earned.",
	"Who a person truly is cannot be seen with the eye.",
	"Understanding is honoring the truth beneath the surface.",
	"Who's the more foolish, the fool or the fool who follows him?",
	"The first step toward loyalty is trust.",
	"The path of ignorance is guided by fear.",
	"The wise man leads, the strong man follows.",
	"Our actions define our legacy.",
	"Where we are going always reflects where we came from.",
	"Those who enslave others inevitably become slaves themselves.",
	"Great hope can come from small sacrifices.",
	"Friendship shows us who we really are.",
	"All warfare is based on deception.",
	"Keep your friends close, but keep your enemies closer.",
	"The strong survive, the noble overcome.",
	"Trust is the greatest of gifts, but it must be earned.",
	"One must let go of the past to hold onto the future.",
	"Who we are never changes, who we think we are does.",
	"A fallen enemy may rise again, but the reconciled one is truly vanquished.",
	"The enemy of my enemy is my friend.",
	"Strength in character can defeat strength in numbers.",
	"Fear is a malleable weapon.",
	"To seek something is to believe in its possibility.",
	"Struggles often begin and end with the truth.",
	"Disobedience is a demand for change.",
	"He who faces himself, finds himself.",
	"The young are often underestimated.",
	"When we rescue others, we rescue ourselves.",
	"Choose your enemies wisely, as they may be your last hope.",
	"Humility is the only defense against humiliation.",
	"When all seems hopeless, a true hero gives hope.",
	"A soldier's most powerful weapon is courage.",
	"You must trust in others or success is impossible.",
	"One vision can have many interpretations.",
	"Alliances can stall true intentions.",
	"Morality separates heroes from villains.",
	"Sometimes even the smallest doubt can shake the greatest belief.",
	"Courage begins by trusting oneself.",
	"Never become desperate enough to trust the untrustworthy.",
	"Never give up hope, no matter how dark things seem.",
}

type Moral struct {
	Message string `json:"message"`
	Key     int `json:"moral_id"`
}

type ErrorMessage struct {
	Error string
}

func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept", "application/vnd.morals."+appVersionStr+"+json")
		w.Header().Set("Content-Type", "application/json")
		fn(w, r)
	}
}

func logHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		fn(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
}

func MoralHandler(w http.ResponseWriter, r *http.Request) {
  rand.Seed(time.Now().Unix())
	key   := rand.Intn(len(morals) - 1) + 1
	moral := &Moral{Message: morals[key], Key: key}
	resp, _ := json.Marshal(moral)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func NewMoralHandler(w http.ResponseWriter, r *http.Request) {
	var moral Moral
	payload, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(payload, &moral)
	resp, _ := json.Marshal(&Moral{Message: moral.Message, Key: moral.Key})
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal(&Moral{})
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func portNumber() string {
	port := os.Getenv("PORT")
	if port == "" {
		port  = "5000"
	}
  return port
}

func init(){
	// var err error
	// // db, err = gorm.Open("postgres", "user=readerwriter dbname=postgres sslmode=disable")
	// if err != nil {
	// 	panic(err)
	// 	return
	// }
}

func main() {
	m := pat.New()
	m.Get("/morals", commonHeaders(logHandler(MoralHandler)))
	m.Post("/morals", commonHeaders(logHandler(NewMoralHandler)))
	m.Get("/", commonHeaders(logHandler(IndexHandler)))
	http.Handle("/", m)
	http.ListenAndServe(":"+portNumber(), nil)
}
