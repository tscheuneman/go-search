type EventHandler = (...passthroughParams: any) => void;

class EventBus {
  eventMap: Record<string, EventHandler[]>;

  constructor() {
    this.eventMap = {};
  }

  subscribe(name: string, handler: EventHandler) {
    if (name && handler) {
      this.eventMap[name] = this.eventMap[name] || [];
      this.eventMap[name].push(handler);
    } else {
      throw new Error("Can't subscript to event");
    }
  }

  remove(name: string, handler?: EventHandler) {
    if(!handler) {
      delete this.eventMap[name];
    }
    if(this.eventMap[name]) {
      const handlerIndex = this.eventMap[name].indexOf(handler as EventHandler);
      if(handlerIndex !== -1) {
        this.eventMap[name].splice(handlerIndex, 1);
      }
    }
  }

  trigger(name: string, ...passthroughParams: any) {
    if(this.eventMap[name]) {
      this.eventMap[name].forEach((handler) => {
        handler(...passthroughParams);
      });
    }
  }

  reset() {
    this.eventMap = {};
  }


}

export default new EventBus();