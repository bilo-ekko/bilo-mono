import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { EquivalentsModule } from './equivalents/equivalents.module';
import { QuotesModule } from './quotes/quotes.module';

@Module({
  imports: [EquivalentsModule, QuotesModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
