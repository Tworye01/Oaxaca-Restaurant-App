import { Card } from "../types/card";

export function validateCard(card: Card): boolean {
  if (!/^(?:4[0-9]{12}(?:[0-9]{3})?)$/.test(card.cardNumber)) {
    return false;
  }

  if (card.cardNumber.length != 16) {
    return false;
  }

  if (card.securityNumber.length != 3) {
    return false;
  }

  if (card.expirationDate < new Date()) {
    return false;
  }

  if (
    card.cardName === "" ||
    card.cardName == undefined ||
    card.cardName == null
  ) {
    return false;
  }

  return true;
}
