export type APIPriceResponse = {
  price: number;
};

export type APILatestPricesResponse = {
  prices: {
    price: number;
    startDate: string;
    endDate: string;
  }[];
};

const date = new Date();
const today = `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, "0")}-${date.getDate()}`;
const currentHour = `${date.getHours().toString().padStart(2, "0")}`;
