

$(document).ready(function() {


    // Delete link request
    $('#example').on('click','.deletelink', function (e) {
 
		e.preventDefault();

        if (confirm('Επιθυμείτε να διαγραφεί αυτή η εγγραφή?')){
	        
			var targetUrl = $(this).attr("href");
			
			$.ajax({
		    	url: targetUrl ,
			    type: 'DELETE',
			    success: function(result) {
			        // Refresh page
			        location.reload();
			    }
			});

			 $('.alert-delete').show(); 

			 $('.alert-delete').delay(6000).slideUp(200,function(){
	
	                   $(this).alert('close');
                    })			
	    }	
	});
    
    // View link request
	$('#example').on('click','.viewlink', function (e) {
		
		e.preventDefault();
        
	    var targetUrl = $(this).attr("href");
			
			$.ajax({
		    	url: targetUrl,
			    type: 'GET',
			    success: function(result) {
			        
			        $('#title_view').val(result.title);
                    $('#description_view').val(result.description);
			    },
			    dataType: "json"
			});	
    	
	}); 
  

    // Insert link request
	$(".insertlink").on("click", function(e){

		    e.preventDefault();
                
			var targetUrl = $(this).attr("href");
				
			$.ajax({
		    	url: targetUrl ,
			    type: 'POST',
			    data: JSON.stringify({
                      "title":$('#title_insert').val(),
			          "description":$('#description_insert').val()
			    }),
			    success: function(result) {
			        // Refresh page
			        location.reload();
                                       
			    },
			     dataType: "json"
			});

				$('.alert-insert').show();

				$('.alert-insert').delay(6000).slideUp(200,function(){
	
	                $(this).alert('close');
                 })
	    	
	});

    // Edit link request
    $('#example').on('click','.editlink', function (e) {
		
		e.preventDefault();
        
	    var targetUrl = $(this).attr("href");
			
			$.ajax({
		    	url: targetUrl ,
			    type: 'GET',
			    success: function(result) {
			        
			        $('#title_update').val(result.title);
                    $('#description_update').val(result.description);
                    $('.updatelink').attr("href",targetUrl);

			    },
			    dataType: "json"
			});			
	    	
	});
   
   // Update link request 
   $(".updatelink").on("click", function(e){

		   e.preventDefault();
                
			var targetUrl = $(this).attr("href");
			
			$.ajax({
		    	url: targetUrl ,
			    type: 'PUT',
			    data: JSON.stringify({
                      "title":$('#title_update').val(),
			          "description":$('#description_update').val()
			    }),
			    success: function(result) {
			        // Refresh page
			        location.reload(); 

			    },
			     dataType: "json"
			});	

			 $('.alert-update').show();

			 $('.alert-update').delay(6000).slideUp(200,function(){
	
	                $(this).alert('close');
                })		
	    	
	    });
   
   // alert messages initialization
  $('.alert-insert').hide();
  $('.alert-update').hide();
  $('.alert-delete').hide();

   
});