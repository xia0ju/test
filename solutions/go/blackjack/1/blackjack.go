package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
    case "ace":
        return 11
    case "two":
        return 2
    case "three":
        return 3
    case "four":
        return 4
   	case "five":
        return 5
	case "six":
        return 6
    case "seven":
        return 7
    case "eight":
        return 8
    case "nine":
        return 9
    case "ten","jack","queen","king":
        return 10
    default:
        return 0
    
    
	    
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
  
    var d=ParseCard(dealerCard)
    var c1=ParseCard(card1)
    var c2=ParseCard(card2)
     var s=c1+c2
    if(card1=="ace"&&card2=="ace"){
        return "P"
    }
    if s==21{
        if dealerCard!="ace"&&dealerCard!="jack"&&dealerCard!="queen"&&dealerCard!="king"&&dealerCard!="ten"{
             return "W"
        }
        return "S"
    }
    if s>=17 &&s<=20{
        return "S"
    }
    if s>=12 &&s<=16{
        if d>=7{
            return "H"
        }
        return "S"
    }
    if s<=11{
        return "H"
    }
    return "H"
    

    
    
}
