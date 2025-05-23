/*
 * This service represents the chat state on the client side.
 * It's responsible for interactions with the server app.
 * Refer to Message.go on how WebScoket messages are formatted
 */

import { Injectable } from '@angular/core';
import { UserProfileService } from './user-profile.service';

import { webSocket, WebSocketSubject } from 'rxjs/webSocket';

@Injectable({
  providedIn: 'root'
})
export class ChannelService {

  m_ProfileService: UserProfileService;
  m_Messages:       Message[];
  m_Lobby:          LobbyEntry[];
  m_Socket!:        WebSocketSubject<unknown>;

  constructor(profileService: UserProfileService) {

    this.m_ProfileService   = profileService;
    this.m_Messages         = new Array<Message>();
    this.m_Lobby            = new Array<LobbyEntry>();

    this.Connect("ws://localhost:8000"); //TODO: this should not be hard coded, use .env or database

    this.m_Socket.subscribe(this.OnMessageReceived.bind(this));

   }

   OnMessageReceived<T>(message: T | any): void {

    switch (message.type) {
      
      case MessageType.UserConnectMSG:
        this.m_Lobby.push(new LobbyEntry(message.sender, message.color));
        break;

      case MessageType.UserDisconnectMSG:
        this.m_Lobby.splice(this.m_Lobby.findIndex(target => target.m_User === message.sender), 1);
        break;

      case MessageType.TextMSG:
        this.m_Messages.push(new Message(message.sender, message.color, message.text));
        break;

    }
    
  }

   Send(text: string): void {

    let message = {
      type:     MessageType.TextMSG,
      sender:   this.m_ProfileService.m_Username,
      color:    this.m_ProfileService.m_Color,
      text:     text
    }
    
    this.m_Socket.next(message);

   }

   Connect(url: string): void {
     
     this.m_Socket = webSocket(url);
     
     let message = {
      type:     MessageType.UserConnectMSG,
      sender:   this.m_ProfileService.m_Username,
      color:    this.m_ProfileService.m_Color,
      text:     ""
    }

    this.m_Socket.next(message);

   }

}

class LobbyEntry {

  m_User:   string;
  m_Color:  string;

  constructor(user: string, color: string) {

    this.m_User   = user;
    this.m_Color  = color

  }

}

class Message {

  m_Sender: string;
  m_Color:  string;
  m_Text:   string;

  constructor(sender: string, color: string, text: string) {

    this.m_Sender   = sender;
    this.m_Color    = color;
    this.m_Text     = text;

  }

}

enum MessageType {

  UserConnectMSG    = 0,
  UserDisconnectMSG = 1,
  TextMSG           = 2

}
