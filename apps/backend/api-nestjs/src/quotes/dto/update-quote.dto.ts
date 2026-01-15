import { QuoteStatus, QuoteItem } from '../entities/quote.entity';

export class UpdateQuoteDto {
  customerId?: string;
  items?: QuoteItem[];
  currency?: string;
  validUntil?: Date;
  status?: QuoteStatus;
}
