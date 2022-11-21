function Prompt(){
    let toast = function(c){
        const {
            msg = "",
            icon = "success",
            position = "top",
        } = c;

        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})

    }

    let success = function(c){
        const {
            msg = '',
            title = '',
            footer = '',
        } = c;
        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer,
        })
    }

    let error =async function(c){
        const {
            msg = '',
            title = '',
            footer = '',
        } = c;
        await Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer,
        })
    }

    async function custom(c){
        const {
            msg = '',
            title = '',
        } = c;
        const { value: formValues } = await Swal.fire({
            title: title,
            html: msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            willOpen: () =>{
                document.getElementById('desired_date').valueAsDate = new Date()
            },
            preConfirm: () => {
                return [
                    document.getElementById('desired_date').value,
                    document.getElementById('time_of_day').value,
                    document.getElementById('calorie').value,
                    document.getElementById('protein').value,
                    document.getElementById('carbs').value,
                    document.getElementById('fats').value,
                ]
            },

        })
        if (formValues){
            if (formValues.dismiss !== Swal.DismissReason.cancel){
                if (formValues.every(value => value !== "")){
                    if (c.callback !== undefined ){
                        c.callback(formValues)
                        return toast({msg:"Successfully updated stats!"})
                    }
                }else {
                    await error({
                        title: "Not Valid",
                        msg: "Please Fill In All Fields"
                    })
                    return custom(c)
                }
            }
        }
    }

    return {
        toast:toast,
        success: success,
        error: error,
        custom: custom,
    }
}
