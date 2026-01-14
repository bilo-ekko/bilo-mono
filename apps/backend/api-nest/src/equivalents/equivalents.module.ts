import { Module } from '@nestjs/common';
import { EquivalentsService } from './equivalents.service';
import { EquivalentsController } from './equivalents.controller';

@Module({
  controllers: [EquivalentsController],
  providers: [EquivalentsService],
  exports: [EquivalentsService],
})
export class EquivalentsModule {}
