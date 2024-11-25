public class Blackjack {

    public int parseCard(String card) {
        if (card.equals("ace"))
            return 11;
        else if (card.equals("king") || card.equals("queen") || card.equals("jack"))
            return 10;
        else if (card.equals("one"))
            return 1;
        else if (card.equals("two"))
            return 2;
        else if (card.equals("three"))
            return 3;
        else if (card.equals("four"))
            return 4;
        else if (card.equals("five"))
            return 5;
        else if (card.equals("six"))
            return 6;
        else if (card.equals("seven"))
            return 7;
        else if (card.equals("eight"))
            return 8;
        else if (card.equals("nine"))
            return 9;
        else if (card.equals("ten"))
            return 10;
        else 
            return 0;
    }

    public boolean isBlackjack(String card1, String card2) {
        int value1 = this.parseCard(card1);
        int value2 = this.parseCard(card2);
        return (value1 == 10 && value2 == 11) || (value1 == 11 && value2 == 10);
    }

    public String largeHand(boolean isBlackjack, int dealerScore) {
        if (isBlackjack && dealerScore < 10)
            return "W";
        else if (isBlackjack && dealerScore >= 10)
            return "S";
        else
            return "P";
    }

    public String smallHand(int handScore, int dealerScore) {
        if (handScore >= 17)
            return "S";
        else if (handScore <= 11 || dealerScore >= 7)
            return "H";
        else
            return "S";
    }

    // FirstTurn returns the semi-optimal decision for the first turn, given the cards of the player and the dealer.
    // This function is already implemented and does not need to be edited. It pulls the other functions together in a
    // complete decision tree for the first turn.
    public String firstTurn(String card1, String card2, String dealerCard) {
        int handScore = parseCard(card1) + parseCard(card2);
        int dealerScore = parseCard(dealerCard);

        if (20 < handScore) {
            return largeHand(isBlackjack(card1, card2), dealerScore);
        } else {
            return smallHand(handScore, dealerScore);
        }
    }
}
