import { QuoteItem } from '../entities/quote.entity';

export class CreateQuoteDto {
  customerId: string;
  items: QuoteItem[];
  currency?: string;
  validUntil?: Date;
}
