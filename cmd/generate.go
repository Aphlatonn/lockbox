package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

type passType int

const (
	password passType = iota
	passPhrase
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random password",
	Long:  `Generate a random password with the provided length and characters.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			choosenPassType          passType = 0
			choosenCharacters        []string = []string{}
			passLen                  string   = ""
			passPhraseWordsSeparator string   = "-"
			copyToClipboard          bool     = false
		)

		// what password type the user want ?
		whatType := huh.NewSelect[passType]().
			Title("What would you like to generate?").
			Options(
				huh.NewOption("Password", password),
				huh.NewOption("Pass phrase", passPhrase),
			).
			Description(`what is the type of the password you want to generate default (password)`).
			Value(&choosenPassType)
		if err := whatType.Run(); err != nil {
			fmt.Println(err)
			return
		}

		switch choosenPassType {
		case password:
			// what characters the user want ?
			whatCharecters := huh.NewMultiSelect[string]().
				Title("What characters would you like to use?").
				Options(
					huh.NewOption("Uppercase letters (A-Z)", "ABCDEFGHIJKLMNOPQRSTUVWXYZ").Selected(true),
					huh.NewOption("Lowercase letters (a-z)", "abcdefghijklmnopqrstuvwxyz").Selected(true),
					huh.NewOption("Numbers", "0123456789"),
					huh.NewOption("Special characters", "!@#$%^&*()_+-&"),
				).
				Value(&choosenCharacters)
			if err := whatCharecters.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// what is the password lenght
			whatLen := huh.NewInput().
				Title("What is the lenght of the password?").
				Prompt("? ").
				Placeholder("14").
				CharLimit(3).
				Value(&passLen)
			if err := whatLen.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// get the lengh
			length, err := strconv.Atoi(passLen)
			if err != nil {
				fmt.Println("⚠️ Password lengh must be a number")
				return
			}

			// generate the password
			password := generateRandomPassword(strings.Join(choosenCharacters, ""), length)

			// send the password a copy to clipboard button
			hereIsYourPassword := huh.NewConfirm().
				Title("Here is your password:").
				Description(password).
				Affirmative("Copy!").
				Negative("No.").
				Value(&copyToClipboard)
			if err := hereIsYourPassword.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// copy to clipboard if the user want
			if copyToClipboard {
				if err := clipboard.WriteAll(password); err != nil {
					fmt.Println(err)
					return
				}
			}

		case passPhrase:
			// what is the passphrase words number
			whatLen := huh.NewInput().
				Title("What Number of words you want?").
				Prompt("? ").
				Placeholder("6").
				CharLimit(2).
				Value(&passLen)
			if err := whatLen.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// what is the password lenght
			whatSeparator := huh.NewInput().
				Title("What Words separator you want to use?").
				Prompt("? ").
				Placeholder("-").
				CharLimit(1).
				Suggestions([]string{"-", "_"}).
				Value(&passPhraseWordsSeparator)
			if err := whatSeparator.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// get the length
			length, err := strconv.Atoi(passLen)
			if err != nil {
				fmt.Println("⚠️ Passphrase length must be a number")
				return
			}

			// generate the passphrase
			passPhrase := generateRandomPassphrase(length, passPhraseWordsSeparator)

			// send the passphrase a copy to clipboard button
			hereIsYourPassphrase := huh.NewConfirm().
				Title("Here is your passphrase:").
				Description(passPhrase).
				Affirmative("Copy!").
				Negative("No.").
				Value(&copyToClipboard)
			if err := hereIsYourPassphrase.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// copy to clipboard if the user want
			if copyToClipboard {
				if err := clipboard.WriteAll(passPhrase); err != nil {
					fmt.Println(err)
					return
				}
			}
		}

	},
}

// generates a random password using the selected characters and length.
func generateRandomPassword(characterSet string, length int) string {
	var passphrase strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(characterSet))
		passphrase.WriteByte(characterSet[randomIndex])
	}
	return passphrase.String()
}

// generates a random passphrase using the selected words number and separetor.
func generateRandomPassphrase(wordsNumber int, separator string) string {
	var passphrase strings.Builder
	for i := 0; i < wordsNumber; i++ {
		randomIndex := rand.Intn(len(words))
		passphrase.WriteString(words[randomIndex])
		if i != wordsNumber-1 {
			passphrase.WriteString(separator)
		}
	}
	return passphrase.String()
}

var words = []string{
	"apple", "banana", "orange", "grape", "watermelon", "pineapple", "blueberry", "strawberry", "peach", "mango",
	"kiwi", "papaya", "guava", "pear", "plum", "cherry", "apricot", "dragonfruit", "lychee", "pomegranate",
	"raspberry", "blackberry", "cantaloupe", "honeydew", "avocado", "tomato", "fig", "date", "coconut", "jackfruit",
	"carrot", "broccoli", "spinach", "kale", "potato", "onion", "garlic", "pepper", "zucchini", "eggplant",
	"lettuce", "celery", "asparagus", "mushroom", "beet", "pumpkin", "squash", "radish", "turnip", "parsnip",
	"barley", "rice", "quinoa", "oat", "wheat", "corn", "millet", "sorghum", "rye", "spelt",
	"chicken", "beef", "pork", "lamb", "turkey", "duck", "fish", "shrimp", "crab", "lobster",
	"egg", "milk", "cheese", "yogurt", "butter", "cream", "ice cream", "chocolate", "candy", "cookie",
	"cake", "pie", "brownie", "doughnut", "pudding", "jelly", "jam", "syrup", "honey", "sugar",
	"salt", "pepper", "vinegar", "oil", "sauce", "spice", "herb", "nut", "seed", "bean",
	"art", "music", "dance", "theater", "film", "literature", "poetry", "painting", "sculpture", "photography",
	"science", "biology", "chemistry", "physics", "mathematics", "astronomy", "geology", "psychology", "sociology", "anthropology",
	"technology", "computer", "internet", "software", "hardware", "network", "database", "algorithm", "programming", "coding",
	"business", "economics", "finance", "marketing", "management", "strategy", "entrepreneur", "startup", "investment", "portfolio",
	"travel", "adventure", "vacation", "holiday", "tour", "journey", "exploration", "expedition", "cruise", "roadtrip",
	"city", "village", "town", "countryside", "mountain", "river", "lake", "ocean", "beach", "desert",
	"sky", "sun", "moon", "star", "planet", "galaxy", "universe", "atmosphere", "weather", "climate",
	"family", "friend", "community", "society", "culture", "tradition", "custom", "festival", "holiday", "celebration",
	"love", "happiness", "joy", "peace", "friendship", "kindness", "compassion", "empathy", "trust", "loyalty",
	"life", "death", "birth", "growth", "change", "freedom", "justice", "equality", "rights", "responsibility",
	"school", "education", "learning", "knowledge", "wisdom", "teacher", "student", "classroom", "lecture", "homework",
	"university", "college", "degree", "course", "subject", "research", "experiment", "study", "project", "assignment",
	"computer", "laptop", "tablet", "phone", "device", "screen", "keyboard", "mouse", "software", "application",
	"website", "blog", "forum", "social media", "networking", "communication", "chat", "video", "audio", "streaming",
	"news", "information", "data", "analytics", "report", "survey", "feedback", "review", "rating", "evaluation",
	"challenge", "problem", "solution", "strategy", "plan", "goal", "achievement", "success", "failure", "lesson",
	"activity", "exercise", "sport", "game", "competition", "tournament", "match", "team", "player", "coach",
	"work", "job", "career", "profession", "employment", "internship", "position", "salary", "benefits", "vacation",
	"health", "wellness", "fitness", "nutrition", "diet", "exercise", "lifestyle", "mental health", "meditation", "relaxation",
	"art", "craft", "design", "creation", "innovation", "invention", "development", "improvement", "upgrade", "enhancement",
	"challenge", "opportunity", "potential", "talent", "skill", "ability", "capacity", "intelligence", "creativity", "imagination",
	"passion", "interest", "hobby", "leisure", "entertainment", "fun", "relaxation", "recreation", "sports", "games",
	"book", "novel", "story", "tale", "fable", "legend", "myth", "history", "biography", "autobiography",
	"comic", "graphic novel", "magazine", "newspaper", "journal", "article", "essay", "report", "research paper", "thesis",
	"poem", "sonnet", "haiku", "lyric", "verse", "stanza", "rhyme", "meter", "literature", "prose",
	"philosophy", "ethics", "morality", "values", "belief", "faith", "religion", "spirituality", "existence", "purpose",
	"thought", "idea", "concept", "theory", "hypothesis", "principle", "law", "rule", "norm", "standard",
	"language", "communication", "dialogue", "conversation", "discussion", "debate", "argument", "persuasion", "rhetoric", "narrative",
	"culture", "society", "community", "tradition", "heritage", "custom", "belief", "value", "norm", "rule",
	"celebration", "festival", "event", "gathering", "meeting", "conference", "symposium", "workshop", "seminar", "webinar",
	"ceremony", "ritual", "custom", "tradition", "holiday", "feast", "banquet", "party", "reunion", "get-together",
	"memory", "experience", "story", "narrative", "reflection", "journey", "adventure", "quest", "mission", "goal",
	"skill", "talent", "ability", "competence", "expertise", "knowledge", "wisdom", "understanding", "insight", "intuition",
	"action", "reaction", "response", "decision", "choice", "option", "alternative", "strategy", "plan", "tactic",
	"problem", "challenge", "obstacle", "setback", "difficulty", "struggle", "conflict", "dispute", "debate", "argument",
	"solution", "answer", "resolution", "conclusion", "outcome", "result", "success", "failure", "lesson", "experience",
	"journey", "path", "road", "way", "direction", "destination", "goal", "target", "objective", "aspiration",
	"future", "dream", "hope", "wish", "desire", "goal", "vision", "ambition", "motivation", "inspiration",
	"life", "existence", "reality", "truth", "perception", "belief", "knowledge", "understanding", "wisdom", "insight",
	"love", "hate", "joy", "sadness", "fear", "anger", "surprise", "disgust", "trust", "distrust",
	"health", "wellness", "fitness", "nutrition", "exercise", "lifestyle", "mental health", "relaxation", "meditation", "mindfulness",
	"nature", "environment", "ecosystem", "biodiversity", "climate", "weather", "landscape", "habitat", "flora", "fauna",
	"planet", "earth", "universe", "cosmos", "galaxy", "star", "moon", "sun", "sky", "atmosphere",
	"travel", "adventure", "journey", "exploration", "expedition", "discovery", "tour", "trip", "vacation", "holiday",
	"city", "town", "village", "countryside", "mountain", "river", "ocean", "lake", "island", "desert",
	"food", "drink", "meal", "snack", "dessert", "appetizer", "main course", "side dish", "beverage", "cocktail",
	"recipe", "cuisine", "flavor", "taste", "ingredient", "dish", "cooking", "baking", "grilling", "frying",
	"music", "art", "theater", "dance", "film", "literature", "poetry", "painting", "sculpture", "photography",
	"science", "technology", "engineering", "mathematics", "biology", "chemistry", "physics", "astronomy", "geology", "psychology",
	"history", "politics", "economics", "sociology", "anthropology", "philosophy", "theology", "law", "government", "policy",
	"business", "finance", "marketing", "management", "strategy", "entrepreneurship", "startup", "investment", "economy", "trade",
	"work", "job", "career", "profession", "employment", "internship", "position", "salary", "benefits", "vacation",
	"health", "wellness", "fitness", "nutrition", "exercise", "mental health", "stress", "anxiety", "depression", "happiness",
	"success", "failure", "learning", "growth", "improvement", "development", "progress", "change", "transformation", "adaptation",
	"friendship", "relationship", "love", "affection", "kindness", "compassion", "empathy", "support", "trust", "loyalty",
	"family", "community", "society", "culture", "tradition", "heritage", "identity", "diversity", "inclusion", "equality",
	"freedom", "justice", "rights", "responsibility", "citizenship", "volunteering", "activism", "advocacy", "philanthropy", "charity",
	"environment", "sustainability", "conservation", "biodiversity", "climate change", "pollution", "resource", "ecosystem", "habitat", "wildlife",
	"nature", "landscape", "flora", "fauna", "ocean", "river", "mountain", "forest", "desert", "island",
	"space", "universe", "cosmos", "galaxy", "star", "planet", "moon", "sun", "black hole", "asteroid",
	"technology", "innovation", "research", "development", "engineering", "design", "computer", "software", "hardware", "internet",
	"communication", "networking", "social media", "website", "application", "data", "analytics", "algorithm", "programming", "coding",
	"media", "communication", "journalism", "advertising", "public relations", "marketing", "branding", "campaign", "strategy", "content",
	"performance", "metrics", "analysis", "feedback", "evaluation", "report", "study", "survey", "poll", "research",
	"event", "gathering", "conference", "meeting", "symposium", "forum", "workshop", "seminar", "presentation", "lecture",
	"celebration", "festival", "holiday", "tradition", "ritual", "ceremony", "commemoration", "anniversary", "birthday", "wedding",
	"memory", "experience", "story", "narrative", "history", "legacy", "heritage", "culture", "identity", "community",
	"creativity", "imagination", "inspiration", "artistry", "craftsmanship", "design", "innovation", "development", "expression", "performance",
	"joy", "happiness", "contentment", "satisfaction", "gratitude", "appreciation", "excitement", "enthusiasm", "passion", "drive",
	"intelligence", "wisdom", "knowledge", "understanding", "insight", "awareness", "consciousness", "perception", "belief", "faith",
	"justice", "fairness", "equity", "truth", "integrity", "honesty", "trustworthiness", "reliability", "loyalty", "dependability",
	"freedom", "autonomy", "self-determination", "independence", "liberty", "choice", "decision", "option", "alternative", "path",
	"journey", "quest", "exploration", "discovery", "adventure", "risk", "challenge", "opportunity", "potential", "growth",
	"development", "learning", "education", "knowledge", "skill", "competence", "expertise", "talent", "ability", "capacity",
	"mindfulness", "meditation", "reflection", "self-awareness", "introspection", "self-care", "wellness", "health", "fitness", "nutrition",
	"lifestyle", "balance", "harmony", "sustainability", "conservation", "responsibility", "advocacy", "activism", "engagement", "participation",
	"solidarity", "community", "collaboration", "teamwork", "partnership", "cooperation", "support", "help", "kindness", "compassion",
	"love", "affection", "friendship", "relationship", "bond", "connection", "trust", "loyalty", "respect", "understanding",
	"culture", "tradition", "heritage", "identity", "diversity", "inclusion", "equality", "justice", "freedom", "rights",
	"education", "school", "learning", "knowledge", "wisdom", "teacher", "student", "classroom", "curriculum", "education",
	"career", "job", "employment", "work", "internship", "position", "salary", "benefits", "retirement", "pension",
	"travel", "journey", "exploration", "adventure", "vacation", "holiday", "trip", "tour", "excursion", "getaway",
	"nature", "wildlife", "environment", "ecology", "sustainability", "climate", "weather", "geography", "landscape", "flora",
	"fauna", "ocean", "river", "mountain", "forest", "desert", "island", "valley", "plateau", "hill",
	"technology", "innovation", "science", "research", "development", "engineering", "computing", "networking", "information", "communication",
	"media", "journalism", "advertising", "marketing", "branding", "design", "visual", "graphic", "interactive", "digital",
	"performance", "metrics", "analytics", "evaluation", "assessment", "report", "study", "survey", "feedback", "outcome",
	"event", "gathering", "meeting", "conference", "symposium", "workshop", "seminar", "presentation", "lecture", "discussion",
	"celebration", "festival", "tradition", "holiday", "ceremony", "commemoration", "anniversary", "birthday", "wedding", "farewell",
	"memory", "experience", "story", "narrative", "history", "legacy", "culture", "community", "identity", "belonging",
	"creativity", "imagination", "artistry", "craftsmanship", "design", "innovation", "development", "expression", "performance", "interpretation",
	"joy", "happiness", "contentment", "gratitude", "appreciation", "excitement", "enthusiasm", "passion", "drive", "motivation",
	"intelligence", "wisdom", "knowledge", "understanding", "insight", "awareness", "consciousness", "perception", "belief", "faith",
	"justice", "fairness", "equity", "truth", "integrity", "honesty", "trustworthiness", "reliability", "loyalty", "dependability",
	"freedom", "autonomy", "self-determination", "independence", "liberty", "choice", "decision", "option", "alternative", "path",
	"journey", "quest", "exploration", "discovery", "adventure", "risk", "challenge", "opportunity", "potential", "growth",
	"development", "learning", "education", "knowledge", "skill", "competence", "expertise", "talent", "ability", "capacity",
	"mindfulness", "meditation", "reflection", "self-awareness", "introspection", "self-care", "wellness", "health", "fitness", "nutrition",
	"lifestyle", "balance", "harmony", "sustainability", "conservation", "responsibility", "advocacy", "activism", "engagement", "participation",
	"solidarity", "community", "collaboration", "teamwork", "partnership", "cooperation", "support", "help", "kindness", "compassion",
	"love", "affection", "friendship", "relationship", "bond", "connection", "trust", "loyalty", "respect", "understanding",
	"culture", "tradition", "heritage", "identity", "diversity", "inclusion", "equality", "justice", "freedom", "rights",
	"education", "school", "learning", "knowledge", "wisdom", "teacher", "student", "classroom", "curriculum", "education",
	"career", "job", "employment", "work", "internship", "position", "salary", "benefits", "retirement", "pension",
	"travel", "journey", "exploration", "adventure", "vacation", "holiday", "trip", "tour", "excursion", "getaway",
	"nature", "wildlife", "environment", "ecology", "sustainability", "climate", "weather", "geography", "landscape", "flora",
	"fauna", "ocean", "river", "mountain", "forest", "desert", "island", "valley", "plateau", "hill",
	"technology", "innovation", "science", "research", "development", "engineering", "computing", "networking", "information", "communication",
	"media", "journalism", "advertising", "marketing", "branding", "design", "visual", "graphic", "interactive", "digital",
	"performance", "metrics", "analytics", "evaluation", "assessment", "report", "study", "survey", "feedback", "outcome",
	"event", "gathering", "meeting", "conference", "symposium", "workshop", "seminar", "presentation", "lecture", "discussion",
	"celebration", "festival", "tradition", "holiday", "ceremony", "commemoration", "anniversary", "birthday", "wedding", "farewell",
	"memory", "experience", "story", "narrative", "history", "legacy", "culture", "community", "identity", "belonging",
	"creativity", "imagination", "artistry", "craftsmanship", "design", "innovation", "development", "expression", "performance", "interpretation",
	"joy", "happiness", "contentment", "gratitude", "appreciation", "excitement", "enthusiasm", "passion", "drive", "motivation",
	"intelligence", "wisdom", "knowledge", "understanding", "insight", "awareness", "consciousness", "perception", "belief", "faith",
	"justice", "fairness", "equity", "truth", "integrity", "honesty", "trustworthiness", "reliability", "loyalty", "dependability",
	"freedom", "autonomy", "self-determination", "independence", "liberty", "choice", "decision", "option", "alternative", "path",
	"journey", "quest", "exploration", "discovery", "adventure", "risk", "challenge", "opportunity", "potential", "growth",
	"development", "learning", "education", "knowledge", "skill", "competence", "expertise", "talent", "ability", "capacity",
	"mindfulness", "meditation", "reflection", "self-awareness", "introspection", "self-care", "wellness", "health", "fitness", "nutrition",
	"lifestyle", "balance", "harmony", "sustainability", "conservation", "responsibility", "advocacy", "activism", "engagement", "participation",
	"solidarity", "community", "collaboration", "teamwork", "partnership", "cooperation", "support", "help", "kindness", "compassion",
	"love", "affection", "friendship", "relationship", "bond", "connection", "trust", "loyalty", "respect", "understanding",
	"culture", "tradition", "heritage", "identity", "diversity", "inclusion", "equality", "justice", "freedom", "rights",
	"education", "school", "learning", "knowledge", "wisdom", "teacher", "student", "classroom", "curriculum", "education",
	"career", "job", "employment", "work", "internship", "position", "salary", "benefits", "retirement", "pension",
	"travel", "journey", "exploration", "adventure", "vacation", "holiday", "trip", "tour", "excursion", "getaway",
	"nature", "wildlife", "environment", "ecology", "sustainability", "climate", "weather", "geography", "landscape", "flora",
	"fauna", "ocean", "river", "mountain", "forest", "desert", "island", "valley", "plateau", "hill",
	"technology", "innovation", "science", "research", "development", "engineering", "computing", "networking", "information", "communication",
	"media", "journalism", "advertising", "marketing", "branding", "design", "visual", "graphic", "interactive", "digital",
	"performance", "metrics", "analytics", "evaluation", "assessment", "report", "study", "survey", "feedback", "outcome",
	"event", "gathering", "meeting", "conference", "symposium", "workshop", "seminar", "presentation", "lecture", "discussion",
	"celebration", "festival", "tradition", "holiday", "ceremony", "commemoration", "anniversary", "birthday", "wedding", "farewell",
	"memory", "experience", "story", "narrative", "history", "legacy", "culture", "community", "identity", "belonging",
	"creativity", "imagination", "artistry", "craftsmanship", "design", "innovation", "development", "expression", "performance", "interpretation",
	"joy", "happiness", "contentment", "gratitude", "appreciation", "excitement", "enthusiasm", "passion", "drive", "motivation",
	"intelligence", "wisdom", "knowledge", "understanding", "insight", "awareness", "consciousness", "perception", "belief", "faith",
	"justice", "fairness", "equity", "truth", "integrity", "honesty", "trustworthiness", "reliability", "loyalty", "dependability",
	"freedom", "autonomy", "self-determination", "independence", "liberty", "choice", "decision", "option", "alternative", "path",
	"journey", "quest", "exploration", "discovery", "adventure", "risk", "challenge", "opportunity", "potential", "growth",
	"development", "learning", "education", "knowledge", "skill", "competence", "expertise", "talent", "ability", "capacity",
	"mindfulness", "meditation", "reflection", "self-awareness", "introspection", "self-care", "wellness", "health", "fitness", "nutrition",
	"lifestyle", "balance", "harmony", "sustainability", "conservation", "responsibility", "advocacy", "activism", "engagement", "participation",
	"solidarity", "community", "collaboration", "teamwork", "partnership", "cooperation", "support", "help", "kindness", "compassion",
	"love", "affection", "friendship", "relationship", "bond", "connection", "trust", "loyalty", "respect", "understanding",
	"culture", "tradition", "heritage", "identity", "diversity", "inclusion", "equality", "justice", "freedom", "rights",
	"education", "school", "learning", "knowledge", "wisdom", "teacher", "student", "classroom", "curriculum", "education",
	"career", "job", "employment", "work", "internship", "position", "salary", "benefits", "retirement", "pension",
	"travel", "journey", "exploration", "adventure", "vacation", "holiday", "trip", "tour", "excursion", "getaway",
	"nature", "wildlife", "environment", "ecology", "sustainability", "climate", "weather", "geography", "landscape", "flora",
	"fauna", "ocean", "river", "mountain", "forest", "desert", "island", "valley", "plateau", "hill",
	"technology", "innovation", "science", "research", "development", "engineering", "computing", "networking", "information", "communication",
	"media", "journalism", "advertising", "marketing", "branding", "design", "visual", "graphic", "interactive", "digital",
	"performance", "metrics", "analytics", "evaluation", "assessment", "report", "study", "survey", "feedback", "outcome",
	"event", "gathering", "meeting", "conference", "symposium", "workshop", "seminar", "presentation", "lecture", "discussion",
	"celebration", "festival", "tradition", "holiday", "ceremony", "commemoration", "anniversary", "birthday", "wedding", "farewell",
	"memory", "experience", "story", "narrative", "history", "legacy", "culture", "community", "identity", "belonging",
	"creativity", "imagination", "artistry", "craftsmanship", "design", "innovation", "development", "expression", "performance", "interpretation",
	"joy", "happiness", "contentment", "gratitude", "appreciation", "excitement", "enthusiasm", "passion", "drive", "motivation",
	"intelligence", "wisdom", "knowledge", "understanding", "insight", "awareness", "consciousness", "perception", "belief", "faith",
	"justice", "fairness", "equity", "truth", "integrity", "honesty", "trustworthiness", "reliability", "loyalty", "dependability",
	"freedom", "autonomy", "self-determination", "independence", "liberty", "choice", "decision", "option", "alternative", "path",
	"journey", "quest", "exploration", "discovery", "adventure", "risk", "challenge", "opportunity", "potential", "growth",
	"development", "learning", "education", "knowledge", "skill", "competence", "expertise", "talent", "ability", "capacity",
}
