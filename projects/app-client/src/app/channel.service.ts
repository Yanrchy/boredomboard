import { Injectable } from '@angular/core';
import { UserProfileService } from './user-profile.service';

import { webSocket, WebSocketSubject } from 'rxjs/webSocket';

@Injectable({
  providedIn: 'root'
})
export class ChannelService {

  m_ProfileService: UserProfileService;
  m_Messages:       Message[];
  m_Socket!:        WebSocketSubject<unknown>;

  constructor(profileService: UserProfileService) {

    this.m_ProfileService = profileService;
    this.m_Messages = new Array<Message>();

    this.Connect("ws://localhost:8000"); //TODO: this should not be hard coded, use .env or database

    this.m_Socket.subscribe(this.OnMessageReceived.bind(this));

   }

   OnMessageReceived<T>(message: T | any): void {

    this.m_Messages.push(new Message(message.sender, message.color, message.text));
    
  }

   Send(text: string): void {

    this.m_Messages.push(new Message(this.m_ProfileService.m_Username, this.m_ProfileService.m_Color, text));

   }

   Connect(url: string): void {

    this.m_Socket = webSocket(url);

   }

}

class Message {

  m_Sender: string;
  m_Color:  string;
  m_Text:   string;

  constructor(sender: string, color: string, text: string) {

    this.m_Sender = sender;
    this.m_Color = color;
    this.m_Text = text;

  }

}
