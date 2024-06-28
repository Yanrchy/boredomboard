import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ChannelService {

  m_Messages: string[];

  constructor() {

    this.m_Messages = new Array<string>();

   }

   Send(message: string): void {

    this.m_Messages.push(message);

   }

}
