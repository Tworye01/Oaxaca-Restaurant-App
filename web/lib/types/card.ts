export type Card = {
  cardName: string;
  cardNumber: string;
  expirationDate: Date;
  securityNumber: string;
};

export const getBlankCard = (): Card => {
  return {
    cardName: "",
    cardNumber: "0",
    expirationDate: new Date(),
    securityNumber: "0",
  };
};
