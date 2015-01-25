!function($, Backbone, _) {
  $.fn.serializeObject = function() {
    var o = {};
    var a = this.serializeArray();
    $.each(a, function() {
      if (o[this.name] !== undefined) {
        if (!o[this.name].push) {
          o[this.name] = [o[this.name]];
        }
        o[this.name].push(this.value || '');
      } else {
        o[this.name] = this.value || '';
      }
    });
    return o;
  };

  var App = {
    Views: {},
    Models: {},
    Helpers: {},
    initialize: function() {
      App.router = new App.Router();
      App.session = App.Helpers.Session;
      Backbone.history.start({pushState: true, hashChange: false});
    }
  };

  App.Router = Backbone.Router.extend({
    routes: {
      ""          : "index",
      "dashboard" : "dashboard"
    },
    index: function() {
      App.navigationBarView = new App.Views.NavigationBar();
      App.indexView         = new App.Views.Index();
    },
    dashboard: function() {
      App.navigationBarView = new App.Views.NavigationBar();
      App.dashboardView     = new App.Views.Dashboard();
    }
  });

  App.Views.Index = Backbone.View.extend({ });

  App.Views.SessionModal = Backbone.View.extend({
    events: {
      "submit form" : "sendForm"
    },

    initialize: function() {
      this.errorContainer = this.$el.find('.alert.alert-danger');
    },

    sendForm: function(e) {
      e.preventDefault();
      this.sendData(this.$el.find('form'));
      return false;
    },

    sendData: function($form) {
      var data    = $form.serializeObject()
        , request = $.ajax({ url: $form.attr('action'), type: "POST", data: JSON.stringify({ user: data })})
        , self    = this

      request.success(function(response) {
        localStorage.setItem("authToken", response.token);
        document.location.href = '/dashboard';
      });

      request.error(function(response) {
        self.errorContainer.html(response.responseJSON.error);
        self.errorContainer.show();
      });
    }
  });

  App.Views.NavigationBar = Backbone.View.extend({
    el: "header.js-navigation-bar",

    events: {
      "click .js-open-sign-in-modal" : "openSignInModal",
      "click .js-open-sign-up-modal" : "openSignUpModal",
      "click .js-sign-out-link"      : "signOut"
    },

    signInModal: this.$('#signInModal'),
    signUpModal: this.$('#signUpModal'),

    initialize: function() {
      if (App.session.authenticated()) {
        this.hideLoginLinks();
      } else {
        App.signInView = new App.Views.SessionModal({ el: "#signInModal" });
        App.signUpView = new App.Views.SessionModal({ el: "#signUpModal" });
      }
    },

    openSignInModal: function(e) {
      this.signInModal.modal('show');
    },

    openSignUpModal: function(e) {
      this.signUpModal.modal('show');
    },

    hideLoginLinks: function() {
      this.$el.find('.js-not-authenticated').hide();
      this.$el.find('.js-authenticated').show();
    },
    signOut: function(e) {
      e.preventDefault()
      $.ajax({
        url: e.currentTarget.href + "?token=" + App.session.token(),
        type: 'DELETE'
      }).success(function(response){
        document.location.href = '/';
        localStorage.removeItem('authToken');
      });
      return false;
    }
  });

  App.Views.Dashboard = Backbone.View.extend({ });

  App.Helpers.Session = {
    authenticated: function() {
      return localStorage.getItem("authToken") != null;
    },
    token: function() {
      return localStorage.getItem('authToken');
    }
  };

  $(document).ready(function() {
    App.initialize();
  });
}(jQuery, Backbone, _);
