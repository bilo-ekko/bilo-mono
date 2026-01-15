export class Quote {
  id: string;
  customerId: string;
  items: QuoteItem[];
  totalAmount: number;
  currency: string;
  status: QuoteStatus;
  validUntil: Date;
  createdAt: Date;
  updatedAt: Date;
}

export interface QuoteItem {
  productId: string;
  quantity: number;
  unitPrice: number;
  totalPrice: number;
}

export enum QuoteStatus {
  DRAFT = 'draft',
  PENDING = 'pending',
  ACCEPTED = 'accepted',
  REJECTED = 'rejected',
  EXPIRED = 'expired',
}
