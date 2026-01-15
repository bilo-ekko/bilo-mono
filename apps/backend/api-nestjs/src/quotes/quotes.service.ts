import { Injectable, NotFoundException } from '@nestjs/common';
import { CreateQuoteDto } from './dto/create-quote.dto';
import { UpdateQuoteDto } from './dto/update-quote.dto';
import { Quote, QuoteStatus } from './entities/quote.entity';

@Injectable()
export class QuotesService {
  private quotes: Map<string, Quote> = new Map();
  private idCounter = 1;

  constructor() {
    // Seed with sample data
    this.seed();
  }

  private seed() {
    const sampleQuotes = [
      {
        customerId: 'customer-1',
        items: [
          {
            productId: 'carbon-offset-1',
            quantity: 100,
            unitPrice: 15,
            totalPrice: 1500,
          },
          {
            productId: 'tree-planting-1',
            quantity: 50,
            unitPrice: 25,
            totalPrice: 1250,
          },
        ],
        currency: 'USD',
      },
      {
        customerId: 'customer-2',
        items: [
          {
            productId: 'solar-energy-1',
            quantity: 200,
            unitPrice: 30,
            totalPrice: 6000,
          },
        ],
        currency: 'USD',
      },
    ];

    sampleQuotes.forEach((data) => {
      const id = String(this.idCounter++);
      const totalAmount = data.items.reduce(
        (sum, item) => sum + item.totalPrice,
        0,
      );
      const validUntil = new Date();
      validUntil.setDate(validUntil.getDate() + 30); // Valid for 30 days

      this.quotes.set(id, {
        id,
        customerId: data.customerId,
        items: data.items,
        totalAmount,
        currency: data.currency,
        status: QuoteStatus.PENDING,
        validUntil,
        createdAt: new Date(),
        updatedAt: new Date(),
      });
    });
  }

  create(createQuoteDto: CreateQuoteDto): Quote {
    const id = String(this.idCounter++);
    const totalAmount = createQuoteDto.items.reduce(
      (sum, item) => sum + item.totalPrice,
      0,
    );

    const validUntil = createQuoteDto.validUntil || new Date();
    if (!createQuoteDto.validUntil) {
      validUntil.setDate(validUntil.getDate() + 30); // Default 30 days validity
    }

    const quote: Quote = {
      id,
      customerId: createQuoteDto.customerId,
      items: createQuoteDto.items,
      totalAmount,
      currency: createQuoteDto.currency || 'USD',
      status: QuoteStatus.DRAFT,
      validUntil,
      createdAt: new Date(),
      updatedAt: new Date(),
    };

    this.quotes.set(id, quote);
    return quote;
  }

  findAll(): Quote[] {
    return Array.from(this.quotes.values());
  }

  findOne(id: string): Quote {
    const quote = this.quotes.get(id);
    if (!quote) {
      throw new NotFoundException(`Quote with ID ${id} not found`);
    }
    return quote;
  }

  findByCustomer(customerId: string): Quote[] {
    return Array.from(this.quotes.values()).filter(
      (quote) => quote.customerId === customerId,
    );
  }

  findByStatus(status: QuoteStatus): Quote[] {
    return Array.from(this.quotes.values()).filter(
      (quote) => quote.status === status,
    );
  }

  update(id: string, updateQuoteDto: UpdateQuoteDto): Quote {
    const quote = this.findOne(id);

    // Recalculate total if items changed
    let totalAmount = quote.totalAmount;
    if (updateQuoteDto.items) {
      totalAmount = updateQuoteDto.items.reduce(
        (sum, item) => sum + item.totalPrice,
        0,
      );
    }

    const updated: Quote = {
      ...quote,
      ...updateQuoteDto,
      totalAmount,
      updatedAt: new Date(),
    };

    this.quotes.set(id, updated);
    return updated;
  }

  remove(id: string): void {
    const quote = this.findOne(id);
    this.quotes.delete(id);
  }
}
