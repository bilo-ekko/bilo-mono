import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
  Query,
} from '@nestjs/common';
import { EquivalentsService } from './equivalents.service';
import { CreateEquivalentDto } from './dto/create-equivalent.dto';
import { UpdateEquivalentDto } from './dto/update-equivalent.dto';

@Controller('equivalents')
export class EquivalentsController {
  constructor(private readonly equivalentsService: EquivalentsService) {}

  @Post()
  create(@Body() createEquivalentDto: CreateEquivalentDto) {
    return this.equivalentsService.create(createEquivalentDto);
  }

  @Get()
  findAll(@Query('category') category?: string) {
    if (category) {
      return this.equivalentsService.findByCategory(category);
    }
    return this.equivalentsService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.equivalentsService.findOne(id);
  }

  @Patch(':id')
  update(
    @Param('id') id: string,
    @Body() updateEquivalentDto: UpdateEquivalentDto,
  ) {
    return this.equivalentsService.update(id, updateEquivalentDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    this.equivalentsService.remove(id);
    return { message: `Equivalent with ID ${id} has been deleted` };
  }
}
