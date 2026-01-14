import { Injectable, NotFoundException } from '@nestjs/common';
import { CreateEquivalentDto } from './dto/create-equivalent.dto';
import { UpdateEquivalentDto } from './dto/update-equivalent.dto';
import { Equivalent } from './entities/equivalent.entity';

@Injectable()
export class EquivalentsService {
  private equivalents: Map<string, Equivalent> = new Map();
  private idCounter = 1;

  constructor() {
    // Seed with sample data
    this.seed();
  }

  private seed() {
    const sampleData: Omit<Equivalent, 'id' | 'createdAt' | 'updatedAt'>[] = [
      {
        category: 'transportation',
        value: 4.6,
        unit: 'kg CO2',
        description: 'Average car trip of 10 miles',
      },
      {
        category: 'energy',
        value: 0.5,
        unit: 'kg CO2',
        description: '1 kWh of electricity',
      },
      {
        category: 'lifestyle',
        value: 2.5,
        unit: 'kg CO2',
        description: 'One meal with beef',
      },
    ];

    sampleData.forEach((data) => {
      const id = String(this.idCounter++);
      this.equivalents.set(id, {
        ...data,
        id,
        createdAt: new Date(),
        updatedAt: new Date(),
      });
    });
  }

  create(createEquivalentDto: CreateEquivalentDto): Equivalent {
    const id = String(this.idCounter++);
    const equivalent: Equivalent = {
      ...createEquivalentDto,
      id,
      createdAt: new Date(),
      updatedAt: new Date(),
    };
    this.equivalents.set(id, equivalent);
    return equivalent;
  }

  findAll(): Equivalent[] {
    return Array.from(this.equivalents.values());
  }

  findOne(id: string): Equivalent {
    const equivalent = this.equivalents.get(id);
    if (!equivalent) {
      throw new NotFoundException(`Equivalent with ID ${id} not found`);
    }
    return equivalent;
  }

  findByCategory(category: string): Equivalent[] {
    return Array.from(this.equivalents.values()).filter(
      (eq) => eq.category.toLowerCase() === category.toLowerCase(),
    );
  }

  update(id: string, updateEquivalentDto: UpdateEquivalentDto): Equivalent {
    const equivalent = this.findOne(id);
    const updated = {
      ...equivalent,
      ...updateEquivalentDto,
      updatedAt: new Date(),
    };
    this.equivalents.set(id, updated);
    return updated;
  }

  remove(id: string): void {
    const equivalent = this.findOne(id);
    this.equivalents.delete(id);
  }
}
