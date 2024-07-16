export type APILatestPricesResponse = {
  prices: {
    price: number;
    startDate: string;
    endDate: string;
  }[];
};
